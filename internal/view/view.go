package view

const error_help = `
------------------------------------------
Escriba una acción valida
------------------------------------------
`

const back_table_helper = "\n\n\n    Presione una tecla para regresar..."

const view_action = "    Acción: "

const view_main = `
    Bienvenido.

    Estan disponibles las siguientes acciones:

    - user           Gestión de usuario para acceder a la api
    - command        Gestión a los commandos de la api
    - source         Gestión a los grupos de direcciones
    - logs           Exportar en un archivo
    - exit           Salir
    `

const view_user = `
    Gestión de usuario.

    Un usuario es indispensable para utilizar la api.

    Estan disponibles las siguientes acciones:

    - create         Crear usuario
    - update         Actualizar datos del usuario
    - back           Regresar
    - exit           Salir
    `

const view_user_create_0 = "Ya existe un usuario valido."
const view_user_create_1 = `
    Crear usuario.

    Debe ingresar los siguientes datos o escribir 'cancel' para regresar
`

const view_user_create_1_1 = "    Usuario: "
const view_user_create_1_2 = "    Contraseña: "
const view_user_create_1_3 = "Creando usuario."

const view_user_update_0 = `
    Actualizar los datos del usuario

    Debe ingresar los siguientes datos o escribir 'cancel' para regresar

`

const view_user_update_0_1 = "    Usuario: "
const view_user_update_0_2 = "    Contraseña: "
const view_user_update_0_3 = "Actualizando usuario."

const view_user_update_1 = `
    No existe un usuario valido.
    `

const view_command = `
    Gestión de comandos.

    Estan disponibles las siguientes acciones:

    - create         Crear comando
    - list           Listar comando
    - delete         Eliminar comando
    - back           Regresar
    - exit           Salir
    `

const view_command_create_0 = `
    Crear comando

    Debe ingresar los siguientes datos o escribir 'cancel' para regresar.

    Notas
    1: La ruta y el comando puede contener argumentos y estos deben tener el
    formato: 'run {{.argument1}} {{.argument2}}' o '/admin/camera/{{.argument1}}/view'
    2: El comando enviado a la ruta es multilinea, si desea terminar de escribir
    solo deje la última linea en blanco y luego de enter
    3: Los mensajes de logs debe ser un json, de la forma { 'status_code': 'message' },
    por ejemplo: { '200': ' {{.now}} - Cambio de nombre correcto'}. Se puede agregar argumentos,
    los argumentos disponibles son:
        - now: Tiempo de la ejecución
        - date: Fecha de la ejecución
        - time: Hora de la ejecución
        - ip: Ip/dirección de la ejecución
        - port: Puerto de la ejecución
        - status: Código de respuesta de la ejecución
        - error: El mensaje de error devuelto por la consulta
    Los mensajes de logs es multilinea, si desea terminar de escribir solo deje la última
    linea en blanco y luego de enter
    4: En caso de no agregar un mensaje de log por defecto, el mensaje utilizado será:
        '{{.now}} - Ejecución responde con código {{.status}} desde {{.ip}} en el puerto {{.port}} - Error: {{.error}}'.
`

const view_command_create_0_1 = "    Nombre: "
const view_command_create_0_2 = "    Ruta: "
const view_command_create_0_3 = "    Comando: "
const view_command_create_0_4 = "    Mensajes del log: "
const view_command_create_0_5 = "    Log por defecto: "

const view_command_create_1 = "Guardando comando"

const view_command_create_2 = "No hay datos completos para el registro de un comando nuevo."

const view_command_list_0 = `
    Lista de comandos

`

const view_command_list_1 = "Sin comandos."

const view_command_delete_0 = `
    Eliminar un comando.

    Debe ingresar los siguientes datos o escribir 'cancel' para regresar.

`

const view_command_delete_1 = "    Ingresar ID: "

const view_command_delete_2 = "Eliminando comando"

const view_command_delete_3 = "Comando no encontrado"

const view_source = `
    Gestión a los grupos de direcciones.

    Estan disponibles las siguientes acciones:

    - create         Crear grupo
    - list           Listar grupos
    - delete         Eliminar grupo
    - view           Ver grupo
    - import_file    Importar grupo
    - export_file    Exportar grupo
    - back           Regresar
    - exit           Salir
    `

const view_source_create_0 = `
    Crear de grupo de direcciones.

    Debe ingresar los siguientes datos o escribir 'cancel' para regresar.

`

const view_source_create_0_1 = "    Nombre: "

const view_source_create_0_2 = "    Comentario: "

const view_source_create_1 = "Guardando grupo de direcciones."

const view_source_list_0 = `
    Lista de grupo de direcciones

`

const view_source_list_1 = "Sin grupo de direcciones."

const view_source_delete_0 = `
    Eliminar un grupo de direcciones.

    Debe ingresar los siguientes datos o escribir 'cancel' para regresar.

`

const view_source_delete_1 = "    Ingresar ID: "

const view_source_delete_2 = "Eliminando grupo de direcciones."

const view_source_delete_3 = "Grupo de direcciones no encontrado."

const view_source_view_0 = `
    Ver grupo de direcciones

    Debe ingresar los siguientes datos o escribir 'cancel' para regresar.

`

const view_source_view_1 = "    Ingresar ID: "

const view_source_view_2 = "Grupo de direcciones no encontrado."

const view_source_view_action = `
    Grupos de direcciones: %s

    Estan disponibles las siguientes acciones:

    - add            Agregar dirección en el grupo
    - list           Listar direcciones del grupo
    - delete         Eliminar dirección del grupo
    - run            Ejecutar comando en este grupo
    - run_ip         Ejecutar comando en una dirección de este grupo
    - export_file    Exportar grupo
    - back           Regresar
    - exit           Salir
    `

const view_source_view_action_add_0 = `
    Agregar dirección en %s.

    Debe ingresar los siguientes datos o escribir 'cancel' para regresar.

`

const view_source_view_action_add_0_1 = "    IP/Dirección: "
const view_source_view_action_add_0_2 = "    Puerto: "

const view_source_view_action_add_1 = "Guardando ip/dirección."

const view_source_view_action_list_0 = `
    Lista de ips/direcciones en grupo de direcciones: %s

`

const view_source_view_action_list_1 = `
    Sin ips/direcciones en grupo de direcciones: %s

`

const view_source_view_action_delete_0 = `
    Eliminar un ip/dirección de en grupo de direcciones: %s.

    Debe ingresar los siguientes datos o escribir 'cancel' para regresar.

`

const view_source_view_action_delete_1 = "    Ingresar ID: "

const view_source_view_action_delete_2 = "Eliminando ip/dirección."

const view_source_view_action_delete_3 = "Ip/dirección no encontrada."

//
const view_source_source_export_0 = `
    Exportar grupo de direcciones

    Debe ingresar los siguientes datos o escribir 'cancel' para regresar.

`

const view_source_source_export_1 = "    Ingresar ID: "
const view_source_source_export_2 = "Grupo de direcciones no encontrado."
const view_source_source_export_3 = "Exportando grupo de direcciones en HOME."

const view_source_source_import_0 = `
    Importar grupo de direcciones

    Debe ingresar los siguientes datos o escribir 'cancel' para regresar.

`

const view_source_source_import_1 = "    Ingresar ubicación: "
const view_source_source_import_2 = "Importando grupo de direcciones."
const view_source_source_import_3 = "El archivo no contiene el fomato correcto o esta corrompido"
const view_source_source_import_4 = "Archivo no encontrado o sin permisos necesarios para acceder a este."

const view_source_view_action_run_0 = "Sin usuario para ejecutar el comando"

const view_source_view_action_run_1 = `
    Ejecutar comando en grupo de direcciones: %s.

    Debe ingresar los siguientes datos o escribir 'cancel' para regresar.

`

const view_source_view_action_run_2 = "    Ingresar ID de la ip/dirección: "
const view_source_view_action_run_3 = "Ip/dirección no encontrada."

const view_source_view_action_run_get_data_1 = "    Ingresar ID del comando: "

const view_source_view_action_run_get_data_2 = `
    Se ejecutará el comando: %s.

    Debe ingresar los siguientes datos o escribir 'cancel' para regresar.

    Notas
    1: Al ser un grupo de direcciones, puede utilizar argumentos de enumeración en los argumentos
    del comando, por ejemplo, si esta renombrando un grupo de camaras, y estas tendran los nombres:
    Camera1, Camera2, etc, puede escribir el argumento 'name' como 'Camera{{ .i }}', i empieza en 0.

`

const view_source_view_action_run_get_data_3 = "    Argumentos de la ruta: %s\n"
const view_source_view_action_run_get_data_4 = "    Ingresar %s: "
const view_source_view_action_run_get_data_5 = "    Argumentos del comando\n"

const view_source_view_action_run_get_data_6 = "No se ha encontrado el comando"

const view_source_logs_export = "Exportando logs en HOME."

func ViewText(view_name string) string {
	texts := map[string]string{
		"error_help":                             error_help,
		"back_table_helper":                      back_table_helper,
		"view_action":                            view_action,
		"view_main":                              view_main,
		"view_user":                              view_user,
		"view_user_create_0":                     view_user_create_0,
		"view_user_create_1":                     view_user_create_1,
		"view_user_create_1_1":                   view_user_create_1_1,
		"view_user_create_1_2":                   view_user_create_1_2,
		"view_user_create_1_3":                   view_user_create_1_3,
		"view_user_update_0":                     view_user_update_0,
		"view_user_update_0_1":                   view_user_update_0_1,
		"view_user_update_0_2":                   view_user_update_0_2,
		"view_user_update_0_3":                   view_user_update_0_3,
		"view_user_update_1":                     view_user_update_1,
		"view_command":                           view_command,
		"view_command_create_0":                  view_command_create_0,
		"view_command_create_0_1":                view_command_create_0_1,
		"view_command_create_0_2":                view_command_create_0_2,
		"view_command_create_0_3":                view_command_create_0_3,
		"view_command_create_0_4":                view_command_create_0_4,
		"view_command_create_0_5":                view_command_create_0_5,
		"view_command_create_1":                  view_command_create_1,
		"view_command_create_2":                  view_command_create_2,
		"view_command_list_0":                    view_command_list_0,
		"view_command_list_1":                    view_command_list_1,
		"view_command_delete_0":                  view_command_delete_0,
		"view_command_delete_1":                  view_command_delete_1,
		"view_command_delete_2":                  view_command_delete_2,
		"view_command_delete_3":                  view_command_delete_3,
		"view_source":                            view_source,
		"view_source_create_0":                   view_source_create_0,
		"view_source_create_0_1":                 view_source_create_0_1,
		"view_source_create_0_2":                 view_source_create_0_2,
		"view_source_create_1":                   view_source_create_1,
		"view_source_list_0":                     view_source_list_0,
		"view_source_list_1":                     view_source_list_1,
		"view_source_delete_0":                   view_source_delete_0,
		"view_source_delete_1":                   view_source_delete_1,
		"view_source_delete_2":                   view_source_delete_2,
		"view_source_delete_3":                   view_source_delete_3,
		"view_source_view_0":                     view_source_view_0,
		"view_source_view_1":                     view_source_view_1,
		"view_source_view_2":                     view_source_view_2,
		"view_source_view_action":                view_source_view_action,
		"view_source_view_action_add_0":          view_source_view_action_add_0,
		"view_source_view_action_add_0_1":        view_source_view_action_add_0_1,
		"view_source_view_action_add_0_2":        view_source_view_action_add_0_2,
		"view_source_view_action_add_1":          view_source_view_action_add_1,
		"view_source_view_action_list_0":         view_source_view_action_list_0,
		"view_source_view_action_list_1":         view_source_view_action_list_1,
		"view_source_view_action_delete_0":       view_source_view_action_delete_0,
		"view_source_view_action_delete_1":       view_source_view_action_delete_1,
		"view_source_view_action_delete_2":       view_source_view_action_delete_2,
		"view_source_view_action_delete_3":       view_source_view_action_delete_3,
		"view_source_source_export_0":            view_source_source_export_0,
		"view_source_source_export_1":            view_source_source_export_1,
		"view_source_source_export_2":            view_source_source_export_2,
		"view_source_source_export_3":            view_source_source_export_3,
		"view_source_source_import_0":            view_source_source_import_0,
		"view_source_source_import_1":            view_source_source_import_1,
		"view_source_source_import_2":            view_source_source_import_2,
		"view_source_source_import_3":            view_source_source_import_3,
		"view_source_source_import_4":            view_source_source_import_4,
		"view_source_view_action_run_0":          view_source_view_action_run_0,
		"view_source_view_action_run_1":          view_source_view_action_run_1,
		"view_source_view_action_run_2":          view_source_view_action_run_2,
		"view_source_view_action_run_3":          view_source_view_action_run_3,
		"view_source_view_action_run_get_data_1": view_source_view_action_run_get_data_1,
		"view_source_view_action_run_get_data_2": view_source_view_action_run_get_data_2,
		"view_source_view_action_run_get_data_3": view_source_view_action_run_get_data_3,
		"view_source_view_action_run_get_data_4": view_source_view_action_run_get_data_4,
		"view_source_view_action_run_get_data_5": view_source_view_action_run_get_data_5,
		"view_source_view_action_run_get_data_6": view_source_view_action_run_get_data_6,
		"view_source_logs_export":                view_source_logs_export,
	}
	var val, ok = texts[view_name]
	if ok {
		return val
	} else {
		return ""
	}
}
