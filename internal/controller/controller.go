package controller

import (
	"bufio"
	"fmt"
	"internal/model"
	"internal/view"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/olekukonko/tablewriter"
	"gorm.io/gorm"
)

var DB *gorm.DB

var Source model.Source

var clear map[string]func()

func init() {
	clear = make(map[string]func())
	clear["linux"] = func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["darwin"] = func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
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
	for _, l := range strings.Repeat("/-\\|", times) {
		Clear()
		fmt.Println(fmt.Sprintf("\n    %s   %c", view.ViewText((message)), l))
		time.Sleep(250 * time.Millisecond)
	}
}

func WaitKey() {
	fmt.Println(view.ViewText("back_table_helper"))
	bufio.NewReader(os.Stdin).ReadString('\n')
}

func Input(text_view string) string {
	PrintView(text_view)
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSuffix(text, "\n")
	return text
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
		"import_file": ControllerMain,
		"if":          ControllerMain,
		"export_file": ControllerMain,
		"ef":          ControllerMain,
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
		ParseScourseList(sources, &parse_sources)
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

func ParseScourseList(sources []model.Source, parse_sources *[][]string) {
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
	ControllerBaseText(fmt.Sprintf(view.ViewText("view_source_view_action"), Source.Name), map[string]func(){
		"add":         ControllerSource,
		"a":           ControllerSource,
		"list":        ControllerSource,
		"l":           ControllerSource,
		"delete":      ControllerSource,
		"d":           ControllerSource,
		"run":         ControllerSource,
		"r":           ControllerSource,
		"run_ip":      ControllerSource,
		"ri":          ControllerSource,
		"export_file": ControllerSource,
		"ef":          ControllerSource,
		"back":        ControllerSource,
		"b":           ControllerSource,
	})
}
