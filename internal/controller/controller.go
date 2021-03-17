package controller

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"internal/model"
	"internal/view"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"os/user"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"text/template"
	"time"

	"github.com/olekukonko/tablewriter"
	"gorm.io/gorm"
)

var DB *gorm.DB

var Source model.Source

var CountIp int

var clear map[string]func()

const log_default = "{{.now}} - Ejecución responde con código {{.status}} desde {{.ip}} en el puerto {{.port}} - Error: {{.error}}"

var GOMAXPROCS int = 10

type IpSer struct {
	Ip   string `json:"ip"`
	Port string `json:"port"`
}

type SourceSer struct {
	Name    string  `json:"name"`
	Comment string  `json:"comment"`
	Ips     []IpSer `json:"ips"`
}

func init() {
	clear = make(map[string]func())
	clear["linux"] = func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		err := cmd.Run()
		if err != nil {
			panic(err)
		}
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		err := cmd.Run()
		if err != nil {
			panic(err)
		}
	}
}

func Clear() {
	value, ok := clear[runtime.GOOS]
	if ok {
		value()
	} else {
		panic("No se puede borrar la terminal, esta plataforma no lo permite")
	}
}

func PrintView(view_name string) {
	fmt.Println(view.ViewText(view_name))
}

func Loading(message string, times int) {
	LoadingBase(view.ViewText((message)), times)
}

func LoadingBase(message string, times int) {
	for _, l := range strings.Repeat("/-\\|", times) {
		Clear()
		fmt.Printf("\n    %s   %c", message, l)
		// fmt.Println(fmt.Sprintf("\n    %s   %c", message, l))
		time.Sleep(250 * time.Millisecond)
	}
}

func WaitKey() {
	fmt.Println(view.ViewText("back_table_helper"))
	_, e := bufio.NewReader(os.Stdin).ReadString('\n')
	// Para ignorar el error
	if e != nil {
		_ = e.Error()
	}
}

func Input(text_view string) string {
	return InputBase(view.ViewText(text_view))
}

func InputBase(message string) string {
	fmt.Println(message)
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSuffix(text, "\n")
	return text
}

func PrintProgressBar(total int) {
	CountIp++
	percent := fmt.Sprintf("%.1f", 100*(float64(CountIp)/float64(total)))
	filledLength := int(100 * CountIp / total)
	bar := strings.Repeat("█", filledLength) + strings.Repeat("-", 100-filledLength)
	Clear()
	fmt.Printf("    %s |%s| %s%% %s\n\n", "Progreso:", bar, percent, "Completo")
}

func format(format string, p map[string]string) string {
	b := &bytes.Buffer{}
	err := template.Must(template.New("").Parse(format)).Execute(b, p)
	if err != nil {
		return ""
	}
	return b.String()
}

func ControllerBase(text_view string, actions map[string]func()) {
	actions["exit"] = ControllerExit
	actions["e"] = ControllerExit
	var function func()
	for {
		PrintView(text_view)
		action := Input("view_action")
		action = strings.ToLower(action)
		function = actions[action]
		if function != nil {
			break
		} else {
			PrintView("error_help")
		}
	}
	function()
}

func ControllerBaseText(text_view string, actions map[string]func()) {
	actions["exit"] = ControllerExit
	actions["e"] = ControllerExit
	var function func()
	for {
		fmt.Println(text_view)
		action := Input("view_action")
		action = strings.ToLower(action)
		function = actions[action]
		if function != nil {
			break
		} else {
			PrintView("error_help")
		}
	}
	function()
}

// Controllers

func ControllerExit() {
	os.Exit(0)
}

func Main(db *gorm.DB) {
	DB = db
	ControllerMain()
}

func ControllerMain() {
	Clear()
	ControllerBase("view_main", map[string]func(){
		"user":    ControllerUser,
		"u":       ControllerUser,
		"command": ControllerCommand,
		"c":       ControllerCommand,
		"source":  ControllerSource,
		"s":       ControllerSource,
		"logs":    ControllerLogsExport,
		"l":       ControllerLogsExport,
		"thread":  ControllerThread,
		"t":       ControllerThread,
	})
}

func ControllerUser() {
	Clear()
	ControllerBase("view_user", map[string]func(){
		"create": ControllerUserCreate,
		"c":      ControllerUserCreate,
		"update": ControllerUserUpdate,
		"u":      ControllerUserUpdate,
		"back":   ControllerMain,
		"b":      ControllerMain,
	})
}

func ControllerUserCreate() {
	Clear()
	var count int64
	DB.Model(&model.User{}).Count(&count)
	if count > 0 {
		Loading("view_user_create_0", 2)
	} else {
		var username string
		var password string
		for {
			PrintView("view_user_create_1")
			username = Input("view_user_create_1_1")
			if strings.ToLower(username) == "cancel" {
				ControllerUser()
				return
			}
			password = Input("view_user_create_1_2")
			if strings.ToLower(password) == "cancel" {
				ControllerUser()
				return
			}
			if username != "" && password != "" {
				break
			}
		}
		user := model.User{Username: username, Password: password}
		Loading("view_user_create_1_3", 2)
		DB.Create(&user)
	}
	ControllerUser()
}

func ControllerUserUpdate() {
	Clear()
	var count int64
	DB.Model(&model.User{}).Count(&count)
	if count > 0 {
		PrintView("view_user_update_0")
		username := Input("view_user_update_0_1")
		if strings.ToLower(username) == "cancel" {
			ControllerUser()
			return
		}
		password := Input("view_user_update_0_2")
		if strings.ToLower(password) == "cancel" {
			ControllerUser()
			return
		}
		if username != "" || password != "" {
			var user model.User
			DB.First(&user)
			if username != "" {
				user.Username = username
			}
			if password != "" {
				user.Password = password
			}
			DB.Save(&user)
			Loading("view_user_update_0_3", 2)
		}
	} else {
		Loading("view_user_update_1", 2)
	}
	ControllerUser()
}

func ControllerCommand() {
	Clear()
	ControllerBase("view_command", map[string]func(){
		"create": ControllerCommandCreate,
		"c":      ControllerCommandCreate,
		"list":   ControllerCommandList,
		"l":      ControllerCommandList,
		"delete": ControllerCommandDelete,
		"d":      ControllerCommandDelete,
		"back":   ControllerMain,
		"b":      ControllerMain,
	})
}

func ControllerCommandCreate() {
	Clear()
	PrintView("view_command_create_0")
	name := Input("view_command_create_0_1")
	if strings.ToLower(name) == "cancel" {
		ControllerCommand()
		return
	}
	url_path := Input("view_command_create_0_2")
	if strings.ToLower(url_path) == "cancel" {
		ControllerCommand()
		return
	}
	if !strings.HasPrefix(url_path, "/") {
		url_path = fmt.Sprintf("/%s", url_path)
	}
	PrintView("view_command_create_0_3")
	var value = ""
	var line string
	reader := bufio.NewReader(os.Stdin)
	for {
		line, _ = reader.ReadString('\n')
		line = strings.TrimSuffix(line, "\n")
		if strings.ToLower(line) == "cancel" {
			ControllerCommand()
			return
		}
		if line != "" {
			if value != "" {
				value = fmt.Sprintf("%s\n%s", value, line)
			} else {
				value = line
			}
		} else {
			break
		}
	}
	PrintView("view_command_create_0_4")
	var message_logs = ""
	for {
		line, _ = reader.ReadString('\n')
		line = strings.TrimSuffix(line, "\n")
		if strings.ToLower(line) == "cancel" {
			ControllerCommand()
			return
		}
		if line != "" {
			if message_logs != "" {
				message_logs = fmt.Sprintf("%s\n%s", message_logs, line)
			} else {
				message_logs = line
			}
		} else {
			break
		}
	}
	PrintView("view_command_create_0_5")
	default_message, _ := reader.ReadString('\n')
	default_message = strings.TrimSuffix(default_message, "\n")
	if strings.ToLower(default_message) == "cancel" {
		ControllerCommand()
		return
	}
	if name != "" && url_path != "" && value != "" {
		if !strings.HasPrefix(message_logs, "{") {
			message_logs = fmt.Sprintf("{%s", message_logs)
		}
		if strings.HasSuffix(message_logs, ",") {
			message_logs = message_logs[:len(message_logs)-1]
		}
		if strings.HasSuffix(message_logs, ",}") {
			message_logs = fmt.Sprintf("%s}", message_logs[:len(message_logs)-2])
		}
		if !strings.HasSuffix(message_logs, "}") {
			message_logs = fmt.Sprintf("%s}", message_logs)
		}
		message_logs = strings.NewReplacer("'", "\"").Replace(message_logs)
		command := model.Command{Name: name, Path: url_path, Value: value, MessageLogs: message_logs, DefaultMessage: default_message}
		DB.Create(&command)
		Loading("view_command_create_1", 2)
	} else {
		Loading("view_command_create_2", 2)
	}
	ControllerCommand()
}

func ControllerCommandList() {
	Clear()
	var commands []model.Command
	DB.Find(&commands)
	if len(commands) > 0 {
		PrintView("view_command_list_0")
		var parse_commands [][]string = [][]string{}
		ParseCommandList(commands, &parse_commands)
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"", "Id", "Nombre"})
		table.SetBorder(false)
		table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
		table.SetAlignment(tablewriter.ALIGN_LEFT)
		table.SetCenterSeparator("")
		table.SetColumnSeparator("")
		table.SetRowSeparator("")
		table.SetTablePadding("\t")
		table.AppendBulk(parse_commands)
		table.Render()
		WaitKey()
	} else {
		Loading("view_command_list_1", 2)
	}
	ControllerCommand()
}

func ParseCommandList(commands []model.Command, parse_commands *[][]string) {
	for _, v := range commands {
		*parse_commands = append(*parse_commands, []string{"", strconv.FormatUint(uint64(v.ID), 10), v.Name})
	}
}

func ControllerCommandDelete() {
	Clear()
	var commands []model.Command
	DB.Find(&commands)
	if len(commands) > 0 {
		PrintView("view_command_delete_0")
		search := Input("view_command_delete_1")
		if strings.ToLower(search) == "cancel" {
			ControllerCommand()
			return
		}
		var command model.Command
		DB.First(&command, search)
		if command.ID > 0 {
			DB.Delete(&command)
			Loading("view_command_delete_2", 2)
		} else {
			Loading("view_command_delete_3", 2)
		}
	} else {
		Loading("view_command_list_1", 2)
	}
	ControllerCommand()
}

func ControllerSource() {
	Clear()
	Source = model.Source{}
	ControllerBase("view_source", map[string]func(){
		"create":      ControllerSourceCreate,
		"c":           ControllerSourceCreate,
		"list":        ControllerScourseList,
		"l":           ControllerScourseList,
		"delete":      ControllerSourceDelete,
		"d":           ControllerSourceDelete,
		"view":        ControllerSourceView,
		"v":           ControllerSourceView,
		"import_file": ControllerSourceImport,
		"if":          ControllerSourceImport,
		"export_file": ControllerSourceExport,
		"ef":          ControllerSourceExport,
		"back":        ControllerMain,
		"b":           ControllerMain,
	})
}

func ControllerSourceCreate() {
	Clear()
	PrintView("view_source_create_0")
	var name string
	for {
		name = Input("view_source_create_0_1")
		if strings.ToLower(name) == "cancel" {
			ControllerSource()
			return
		}
		if name != "" {
			break
		}
	}
	comment := Input("view_source_create_0_2")
	if strings.ToLower(comment) == "cancel" {
		ControllerSource()
		return
	}
	source := model.Source{Name: name, Comment: comment}
	DB.Create(&source)
	Loading("view_source_create_1", 2)
	ControllerSource()
}

func ControllerScourseList() {
	Clear()
	var sources []model.Source
	DB.Find(&sources)
	if len(sources) > 0 {
		PrintView("view_source_list_0")
		var parse_sources [][]string = [][]string{}
		ParseSourseList(sources, &parse_sources)
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"", "Id", "Nombre", "Comentario"})
		table.SetBorder(false)
		table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
		table.SetAlignment(tablewriter.ALIGN_LEFT)
		table.SetCenterSeparator("")
		table.SetColumnSeparator("")
		table.SetRowSeparator("")
		table.SetTablePadding("\t")
		table.AppendBulk(parse_sources)
		table.Render()
		WaitKey()
	} else {
		Loading("view_source_list_1", 2)
	}
	ControllerSource()
}

func ParseSourseList(sources []model.Source, parse_sources *[][]string) {
	for _, v := range sources {
		*parse_sources = append(*parse_sources, []string{"", strconv.FormatUint(uint64(v.ID), 10), v.Name, v.Comment})
	}
}

func ControllerSourceDelete() {
	Clear()
	var sources []model.Source
	DB.Find(&sources)
	if len(sources) > 0 {
		PrintView("view_source_delete_0")
		search := Input("view_source_delete_1")
		if strings.ToLower(search) == "cancel" {
			ControllerSource()
			return
		}
		var source model.Source
		DB.First(&source, search)
		if source.ID > 0 {
			DB.Delete(&source)
			Loading("view_source_delete_2", 2)
		} else {
			Loading("view_source_delete_3", 2)
		}
	} else {
		Loading("view_source_list_1", 2)
	}
	ControllerSource()
}

func ControllerSourceView() {
	Clear()
	PrintView("view_source_view_0")
	search := Input("view_source_view_1")
	if strings.ToLower(search) == "cancel" {
		ControllerSource()
		return
	}
	var source model.Source
	DB.First(&source, search)
	if source.ID > 0 {
		Source = source
		ControllerSourceViewAction()
		return
	} else {
		Loading("view_source_view_2", 2)
	}
	ControllerSource()
}

func ControllerSourceViewAction() {
	Clear()
	ControllerSourceViewActionBase()
}

func ControllerSourceViewActionBase() {
	ControllerBaseText(fmt.Sprintf(view.ViewText("view_source_view_action"), Source.Name), map[string]func(){
		"add":         ControllerSourceViewActionAdd,
		"a":           ControllerSourceViewActionAdd,
		"list":        ControllerSourceViewActionList,
		"l":           ControllerSourceViewActionList,
		"delete":      ControllerSourceViewActionDelete,
		"d":           ControllerSourceViewActionDelete,
		"run":         ControllerSourceViewActionRun,
		"r":           ControllerSourceViewActionRun,
		"run_ip":      ControllerSourceViewActionRunIp,
		"ri":          ControllerSourceViewActionRunIp,
		"export_file": ControllerSourceExport,
		"ef":          ControllerSourceExport,
		"back":        ControllerSource,
		"b":           ControllerSource,
	})
}

func ControllerSourceViewActionAdd() {
	Clear()
	fmt.Println(fmt.Sprintf(view.ViewText("view_source_view_action_add_0"), Source.Name))
	var ip, port string
	for {
		ip = Input("view_source_view_action_add_0_1")
		if strings.ToLower(ip) == "cancel" {
			ControllerSourceViewAction()
			return
		}
		port = Input("view_source_view_action_add_0_2")
		if strings.ToLower(port) == "cancel" {
			ControllerSourceViewAction()
			return
		}
		if ip != "" && port != "" {
			break
		}
	}
	ip_el := model.Ip{Ip: ip, Port: port, SourceID: Source.ID}
	DB.Save(&ip_el)
	Loading("view_source_view_action_add_1", 2)
	ControllerSourceViewAction()
}

func ControllerSourceViewActionList() {
	Clear()
	var ips []model.Ip
	DB.Where("source_iD = ?", Source.ID).Find(&ips)
	if len(ips) > 0 {
		fmt.Println(fmt.Sprintf(view.ViewText("view_source_view_action_list_0"), Source.Name))
		var parse_ips [][]string = [][]string{}
		ParseSourseViewActionList(ips, &parse_ips)
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"", "Id", "IP/Dirección", "Puerto"})
		table.SetBorder(false)
		table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
		table.SetAlignment(tablewriter.ALIGN_LEFT)
		table.SetCenterSeparator("")
		table.SetColumnSeparator("")
		table.SetRowSeparator("")
		table.SetTablePadding("\t")
		table.AppendBulk(parse_ips)
		table.Render()
		WaitKey()
	} else {
		LoadingBase(fmt.Sprintf(view.ViewText("view_source_view_action_list_1"), Source.Name), 2)
	}
	ControllerSourceViewAction()
}

func ParseSourseViewActionList(ips []model.Ip, parse_ips *[][]string) {
	for _, v := range ips {
		*parse_ips = append(*parse_ips, []string{"", strconv.FormatUint(uint64(v.ID), 10), v.Ip, v.Port})
	}
}

func ControllerSourceViewActionDelete() {
	Clear()
	var ips []model.Ip
	DB.Where("source_iD = ?", Source.ID).Find(&ips)
	if len(ips) > 0 {
		fmt.Println(fmt.Sprintf(view.ViewText("view_source_view_action_delete_0"), Source.Name))
		search := Input("view_source_view_action_delete_1")
		if strings.ToLower(search) == "cancel" {
			ControllerSourceViewAction()
			return
		}
		var ip model.Ip
		DB.First(&ip, search)
		if ip.ID > 0 {
			DB.Delete(&ip)
			Loading("view_source_view_action_delete_2", 2)
		} else {
			Loading("view_source_view_action_delete_3", 2)
		}
	} else {
		LoadingBase(fmt.Sprintf(view.ViewText("view_source_view_action_list_1"), Source.Name), 2)
	}
	ControllerSourceViewAction()
}

func ControllerSourceViewActionRun() {
	Clear()
	var user model.User
	DB.First(&user)
	if user.ID > 0 {
		fmt.Println(fmt.Sprintf(view.ViewText("view_source_view_action_run_1"), Source.Name))
		path, command, message_logs, default_message := GetDataRunSource()
		if path == "" && command == "" {
			ControllerSourceViewAction()
			return
		}
		var ips []model.Ip
		DB.Where("source_iD = ?", Source.ID).Find(&ips)
		ExecOnIps(&ips, &user, path, command, message_logs, default_message)
	} else {
		Loading("view_source_view_action_run_0", 2)
		ControllerSourceViewAction()
		return
	}
	ControllerSourceViewActionBase()
}

func ControllerSourceViewActionRunIp() {
	Clear()
	var user model.User
	DB.First(&user)
	if user.ID > 0 {
		fmt.Println(fmt.Sprintf(view.ViewText("view_source_view_action_run_1"), Source.Name))
		search := Input("view_source_view_action_delete_1")
		if strings.ToLower(search) == "cancel" {
			ControllerSourceViewAction()
			return
		}
		var ip model.Ip
		DB.First(&ip, search)
		if ip.ID > 0 {
			path, command, message_logs, default_message := GetDataRunSource()
			if path == "" && command == "" {
				ControllerSourceViewAction()
				return
			}
			ips := []model.Ip{ip}
			ExecOnIps(&ips, &user, path, command, message_logs, default_message)
		} else {
			Loading("view_source_view_action_run_3", 2)
			ControllerSourceViewAction()
			return
		}

	} else {
		Loading("view_source_view_action_run_0", 2)
	}
	ControllerSourceViewActionBase()
}

func GetDataRunSource() (string, string, string, string) {
	search := Input("view_source_view_action_run_get_data_1")
	if strings.ToLower(search) == "cancel" {
		ControllerSourceViewAction()
		return "", "", "", ""
	}
	var command model.Command
	DB.First(&command, search)
	if command.ID > 0 {
		fmt.Println(fmt.Sprintf(view.ViewText("view_source_view_action_run_get_data_2"), command.Name))
		re := regexp.MustCompile(`{{\s*[\w\.]+\s*}}`)
		args_path := re.FindAll([]byte(command.Path), -1)
		args_path_dict := make(map[string]string)
		if len(args_path) > 0 {
			fmt.Println(fmt.Sprintf(view.ViewText("view_source_view_action_run_get_data_3"), command.Path))
			for _, arg := range args_path {
				arg := string(arg)
				arg = strings.TrimSpace(arg[3 : len(arg)-2])
				value := InputBase(fmt.Sprintf(view.ViewText("view_source_view_action_run_get_data_4"), arg))
				if strings.ToLower(value) == "cancel" {
					ControllerSourceViewAction()
					return "", "", "", ""
				}
				args_path_dict[arg] = value
			}

		}
		args_value := re.FindAll([]byte(command.Value), -1)
		args_value_dict := make(map[string]string)
		if len(args_value) > 0 {
			PrintView(("view_source_view_action_run_get_data_5"))
			for _, arg := range args_value {
				arg := string(arg)
				arg = strings.TrimSpace(arg[3 : len(arg)-2])
				value := InputBase(fmt.Sprintf(view.ViewText("view_source_view_action_run_get_data_4"), arg))
				if strings.ToLower(value) == "cancel" {
					ControllerSourceViewAction()
					return "", "", "", ""
				}
				args_value_dict[arg] = value
			}

		}
		var default_message string
		if len(command.DefaultMessage) > 0 {
			default_message = command.DefaultMessage
		} else {
			default_message = log_default
		}
		args_path_dict["i"] = "{{.i}}"
		args_value_dict["i"] = "{{.i}}"
		return format(command.Path, args_path_dict), format(command.Value, args_value_dict), command.MessageLogs, default_message

	} else {
		Loading("view_source_view_action_run_get_data_6", 2)
	}
	return "", "", "", ""
}

func ExecRunSource(user *model.User, ip *model.Ip, path string, command string) (string, int) {
	var target_url string
	if strings.HasPrefix(ip.Ip, "http") {
		target_url = ip.Ip + ":" + ip.Port + path
	} else {
		target_url = "http://" + ip.Ip + ":" + ip.Port + path

	}
	client := http.Client{}
	data := []byte(command)
	req, _ := http.NewRequest("POST", target_url, bytes.NewBuffer(data))
	req.SetBasicAuth(user.Username, user.Password)
	resp, err := client.Do(req)
	if err != nil {
		return err.Error(), 500
	} else {
		defer resp.Body.Close()
		return "None", resp.StatusCode
	}
}

func ExecOnIps(ips *[]model.Ip, user *model.User, path string, command string, message_logs string, default_message string) {
	Clear()
	var logs []string = []string{}
	CountIp = -1
	PrintProgressBar(len(*ips))
	runtime.GOMAXPROCS(GOMAXPROCS)
	var wg sync.WaitGroup
	for i, ip := range *ips {
		wg.Add(1)
		go func(i int, ip model.Ip) {
			defer wg.Done()
			command_i := format(command, map[string]string{"$i": strconv.Itoa(i + 1)})
			error, status := ExecRunSource(user, &ip, path, command_i)
			now := time.Now()
			message_logs_map := make(map[string]string)
			err := json.Unmarshal([]byte(message_logs), &message_logs_map)
			if err != nil {
				panic(err)
			}
			message_log, found := message_logs_map[strconv.Itoa(status)]
			if !found {
				message_log = default_message
			}
			log := model.Log{Value: format(message_log, map[string]string{
				"now":    now.Format("2006-01-02T15:04:05"),
				"date":   now.Format("2006-01-02"),
				"time":   now.Format("15:04:05"),
				"error":  error,
				"ip":     ip.Ip,
				"port":   ip.Port,
				"status": strconv.Itoa(status),
			})}
			DB.Create(&log)
			logs = append(logs, log.Value)
			PrintProgressBar(len(*ips))
		}(i, ip)
	}
	wg.Wait()
	for _, log := range logs {
		fmt.Printf("    - %s\n", log)
	}
	fmt.Println("\n    ------------------------------------------")
}

func ControllerSourceImport() {
	Clear()
	PrintView("view_source_source_import_0")
	search := Input("view_source_source_import_1")
	if strings.ToLower(search) == "cancel" {
		ControllerSource()
		return
	}
	inputfile, err := ioutil.ReadFile(search)
	if err != nil {
		Loading("view_source_source_import_4", 2)

	} else {
		var source_ser SourceSer
		json.Unmarshal([]byte(inputfile), &source_ser)
		if len(source_ser.Name) > 0 {
			source := model.Source{Name: source_ser.Name, Comment: source_ser.Comment}
			DB.Create(&source)
			for _, ip := range source_ser.Ips {
				ip_el := model.Ip{Ip: ip.Ip, Port: ip.Port, SourceID: source.ID}
				DB.Save(&ip_el)
			}
			Loading("view_source_source_import_2", 2)
		} else {

			Loading("view_source_source_import_3", 2)
		}
	}
	ControllerSource()
}

func ControllerSourceExport() {
	Clear()
	var to_action bool
	var source model.Source
	var search string
	if Source.ID > 0 {
		to_action = true
		source = Source
	} else {
		to_action = false
		PrintView("view_source_source_export_0")
		search = Input("view_source_source_export_1")
		if strings.ToLower(search) == "cancel" {
			if to_action {
				ControllerSourceViewAction()
			} else {
				ControllerSource()
			}
			return
		}
		DB.First(&source, search)
	}
	if source.ID == 0 {
		Loading("view_source_source_export_2", 2)
		if to_action {
			ControllerSourceViewAction()
		} else {
			ControllerSource()
		}
		return
	}
	source_ser := SourceSer{
		Name:    source.Name,
		Comment: source.Comment,
		Ips:     []IpSer{},
	}
	var ips []model.Ip
	DB.Where("source_iD = ?", source.ID).Find(&ips)
	for _, ip := range ips {
		source_ser.Ips = append(source_ser.Ips, IpSer{
			Ip:   ip.Ip,
			Port: ip.Port,
		})
	}
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}
	file, _ := json.MarshalIndent(source_ser, "", "    ")
	_ = ioutil.WriteFile(usr.HomeDir+"/"+source.Name+"-"+strconv.Itoa(int(source.ID))+".json", file, 0644)
	Loading("view_source_source_export_3", 2)
	if to_action {
		ControllerSourceViewAction()
	} else {
		ControllerSource()
	}
}

func ControllerLogsExport() {
	Clear()
	var logs []model.Log
	DB.Find(&logs)
	now := time.Now()
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}

	file, _ := os.OpenFile(usr.HomeDir+"/logs-mxcam-"+now.Format("2006-01-02T15:04:05")+".txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	datawriter := bufio.NewWriter(file)
	for _, log := range logs {
		_, _ = datawriter.WriteString(log.Value + "\n")
	}
	datawriter.Flush()
	file.Close()
	Loading("view_source_logs_export", 2)
	ControllerMain()
}

func ControllerThread() {
	Clear()
	fmt.Println(fmt.Sprintf(view.ViewText("view_thread_0"), strconv.Itoa(GOMAXPROCS)))
	cant_s := Input("view_thread_1")
	if strings.ToLower(cant_s) == "cancel" {
		ControllerMain()
		return
	}
	var err error
	var cant int
	cant, err = strconv.Atoi(cant_s)
	if err != nil {
		Loading("view_thread_2", 2)
	} else {
		GOMAXPROCS = cant
		Loading("view_thread_3", 2)
	}
	ControllerMain()
}
