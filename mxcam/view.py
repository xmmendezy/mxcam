# -*- coding: utf-8 -*-

__all__ = [
    'error_help',
    'separator_help',
    'back_table_helper',
    'cancel_input_helper',
    'log_default',
    'view_main',
    'view_user', 'view_user_create_0', 'view_user_create_1', 'view_user_create_1_1', 'view_user_create_1_2', 'view_user_create_1_3',
    'view_user_update_0', 'view_user_update_0_1', 'view_user_update_0_2', 'view_user_update_0_3', 'view_user_update_1',
    'view_command', 'view_command_run_0', 'view_command_create_0', 'view_command_create_0_1', 'view_command_create_0_2', 'view_command_create_0_3',
    'view_command_create_0_4', 'view_command_create_0_5' , 'view_command_create_1', 'view_command_create_2', 'view_command_list_0', 'view_command_list_1',
    'view_command_delete_0', 'view_command_delete_1', 'view_command_delete_2', 'view_command_delete_3',
    'view_source', 'view_source_create_0', 'view_source_create_0_1', 'view_source_create_0_2', 'view_source_create_1',
    'view_source_list_0', 'view_source_list_1', 'view_source_delete_0', 'view_source_delete_1', 'view_source_delete_2', 'view_source_delete_3',
    'view_source_view_0', 'view_source_view_1', 'view_source_view_2', 'view_source_view_action',
    'view_source_view_action_add_0', 'view_source_view_action_add_0_1', 'view_source_view_action_add_0_2', 'view_source_view_action_add_1',
    'view_source_view_action_list_0', 'view_source_view_action_list_1', 'view_source_view_action_delete_0', 'view_source_view_action_delete_1',
    'view_source_view_action_delete_2', 'view_source_view_action_delete_3', 'view_source_source_export_0', 'view_source_source_export_1',
    'view_source_source_export_2', 'view_source_source_export_3', 'view_source_source_import_0', 'view_source_source_import_1',
    'view_source_source_import_2', 'view_source_source_import_3', 'view_source_source_import_4',
    'view_source_view_action_run_0', 'view_source_view_action_run_1', 'view_source_view_action_run_2', 'view_source_view_action_run_3',
    'view_source_view_action_run_get_data_1', 'view_source_view_action_run_get_data_2', 'view_source_view_action_run_get_data_3',
    'view_source_view_action_run_get_data_4', 'view_source_view_action_run_get_data_5', 'view_source_view_action_run_get_data_6',
    'view_source_logs_export',
]

action_help = 'Estan disponibles las siguientes acciones:'
separator_help = '-' * len(action_help)
back_help = '- back           Regresar'
exit_help = '- exit           Salir'
back_table_helper = '\n    Presione una tecla para regresar...'
cancel_input_helper = '\n    Cancelado. Presione una tecla para regresar...'
log_default = '{now} - Ejecución responde con código {status} desde {ip} en el puerto {port}'
error_help = f'''
    {separator_help}
    Escriba una acción valida
    {separator_help}
    '''

view_main = f'''
    Bienvenido.

    {action_help}

    - user           Gestión de usuario para acceder a la api
    - command        Gestión a los commandos de la api
    - source         Gestión a los grupos de direcciones
    - logs           Exportar en un archivo
    {exit_help}
    '''


view_user = f'''
    Gestión de usuario.

    Un usuario es indispensable para utilizar la api.

    {action_help}

    - create         Crear usuario
    - update         Actualizar datos del usuario
    {back_help}
    {exit_help}
    '''

view_user_create_0 = 'Ya existe un usuario valido.'

view_user_create_1 = f'''
    Crear usuario.

    Debe ingresar los siguientes datos o escribir 'cancel' para regresar
    '''

view_user_create_1_1 = '    Usuario: '
view_user_create_1_2 = '    Contraseña: '
view_user_create_1_3 = 'Creando usuario.'


view_user_update_0 = f'''
    Actualizar los datos del usuario

    Debe ingresar los siguientes datos o escribir 'cancel' para regresar

'''

view_user_update_0_1 = '    Usuario: '
view_user_update_0_2 = '    Contraseña: '
view_user_update_0_3 = 'Actualizando usuario.'

view_user_update_1 = f'''
    No existe un usuario valido.
    '''

view_command = f'''
    Gestión de comandos.

    {action_help}

    - create         Crear comando
    - list           Listar comando
    - delete         Eliminar comando
    {back_help}
    {exit_help}
    '''

view_command_run_0 = f'''
    Ejecutar comando

'''

view_command_create_0 = '''
    Crear comando

    Debe ingresar los siguientes datos o escribir 'cancel' para regresar.

    Notas
    1: Los argumentos deben tener el formato: 'run {argument1} {argument1}'
    2: El comando enviado a la ruta es multilinea, si desea terminar de escribir
    solo deje la última linea en blanco y luego de enter
    3: Los mensajes de logs debe ser un json, de la forma { 'status_code': 'message' },
    por ejemplo: { '200': ' {now} - Cambio de nombre correcto'}. Se puede agregar argumentos,
    los argumentos disponibles son:
        - now: Tiempo de la ejecución
        - date: Fecha de la ejecución
        - time: Hora de la ejecución
        - ip: Ip/dirección de la ejecución
        - port: Puerto de la ejecución
        - status: Código de respuesta de la ejecución
    Los mensajes de logs es multilinea, si desea terminar de escribir solo deje la última
    linea en blanco y luego de enter
    4: En caso de no agregar un mensaje de log por defecto, el mensaje utilizado será:
        '{now} - Ejecución responde con código {status} desde {ip} en el puerto {port}'.
'''

view_command_create_0_1 = '    Nombre: '
view_command_create_0_2 = '    Ruta: '
view_command_create_0_3 = '    Comando: '
view_command_create_0_4 = '    Mensajes del log: '
view_command_create_0_5 = '    Log por defecto: '


view_command_create_1 = 'Guardando comando'

view_command_create_2 = 'No hay datos completos para el registro de un comando nuevo.'

view_command_list_0 = f'''
    Lista de comando

'''

view_command_list_1 = f'''
    Sin comandos.

'''

view_command_delete_0 = f'''
    Eliminar un comando.

    Debe ingresar los siguientes datos o escribir 'cancel' para regresar.

'''

view_command_delete_1 = '    Ingresar ID: '

view_command_delete_2 = 'Eliminando comando'

view_command_delete_3 = 'Comando no encontrado'

view_source = f'''
    Gestión a los grupos de direcciones.

    {action_help}

    - create         Crear grupo
    - list           Listar grupos
    - delete         Eliminar grupo
    - view           Ver grupo
    - import_file    Importar grupo
    - export_file    Exportar grupo
    {back_help}
    {exit_help}
    '''

view_source_create_0 = '''
    Crear de grupo de direcciones.

    Debe ingresar los siguientes datos o escribir 'cancel' para regresar.

'''

view_source_create_0_1 = '    Nombre: '
view_source_create_0_2 = '    Comentario: '

view_source_create_1 = 'Guardando grupo de direcciones.'

view_source_list_0 = '''
    Lista de grupo de direcciones

'''

view_source_list_1 = '''
    Sin grupo de direcciones.

'''

view_source_delete_0 = '''
    Eliminar un grupo de direcciones.

    Debe ingresar los siguientes datos o escribir 'cancel' para regresar.

'''

view_source_delete_1 = '    Ingresar ID: '

view_source_delete_2 = 'Eliminando grupo de direcciones.'

view_source_delete_3 = 'Grupo de direcciones no encontrado.'

view_source_view_0 = '''
    Ver grupo de direcciones

    Debe ingresar los siguientes datos o escribir 'cancel' para regresar.

'''

view_source_view_1 = '    Ingresar ID: '

view_source_view_2 = 'Grupo de direcciones no encontrado.'

view_source_view_action = f'''
    Grupos de direcciones: {{source}}

    {action_help}

    - add            Agregar dirección en el grupo
    - list           Listar direcciones del grupo
    - delete         Eliminar dirección del grupo
    - run            Ejecutar comando en este grupo
    - run_ip         Ejecutar comando en una dirección de este grupo
    - export_file    Exportar grupo
    {back_help}
    {exit_help}
    '''

view_source_view_action_add_0 = '''
    Agregar dirección en {source}.

    Debe ingresar los siguientes datos o escribir 'cancel' para regresar.

'''

view_source_view_action_add_0_1 = '    IP/Dirección: '
view_source_view_action_add_0_2 = '    Puerto: '

view_source_view_action_add_1 = 'Guardando ip/dirección.'

view_source_view_action_list_0 = '''
    Lista de ips/direcciones en grupo de direcciones: {source}

'''

view_source_view_action_list_1 = '''
    Sin ips/direcciones en grupo de direcciones: {source}

'''

view_source_view_action_delete_0 = '''
    Eliminar un ip/dirección de en grupo de direcciones: {source}.

    Debe ingresar los siguientes datos o escribir 'cancel' para regresar.

'''

view_source_view_action_delete_1 = '    Ingresar ID: '

view_source_view_action_delete_2 = 'Eliminando ip/dirección.'

view_source_view_action_delete_3 = 'Ip/dirección no encontrada.'


view_source_source_export_0 = '''
    Exportar grupo de direcciones

    Debe ingresar los siguientes datos o escribir 'cancel' para regresar.

'''

view_source_source_export_1 = '    Ingresar ID: '
view_source_source_export_2 = 'Grupo de direcciones no encontrado.'
view_source_source_export_3 = 'Exportando grupo de direcciones en HOME.'

view_source_source_import_0 = '''
    Importar grupo de direcciones

    Debe ingresar los siguientes datos o escribir 'cancel' para regresar.

'''

view_source_source_import_1 = '    Ingresar ubicación: '
view_source_source_import_2 = 'Importando grupo de direcciones.'
view_source_source_import_3 = 'El archivo no contiene el fomato correcto o esta corrompido'
view_source_source_import_4 = 'Archivo no encontrado o sin permisos necesarios para acceder a este.'


view_source_view_action_run_0 = 'Sin usuario para ejecutar el comando'

view_source_view_action_run_1 = '''
    Ejecutar comando en grupo de direcciones: {source}.

    Debe ingresar los siguientes datos o escribir 'cancel' para regresar.

'''

view_source_view_action_run_2 = '    Ingresar ID de la ip/dirección: '
view_source_view_action_run_3 = 'Ip/dirección no encontrada.'

view_source_view_action_run_get_data_1 = '    Ingresar ID del comando: '

view_source_view_action_run_get_data_2 = '''
    Se ejecutará el comando: {command}.

    Debe ingresar los siguientes datos o escribir 'cancel' para regresar.

    Notas
    1: Al ser un grupo de direcciones, puede utilizar argumentos de enumeración en los argumentos
    del comando, por ejemplo, si esta renombrando un grupo de camaras, y estas tendran los nombres:
    Camera1, Camera2, etc, puede escribir el argumento 'name' como 'Camera{{i + 1}}', i empieza en 0.

'''

view_source_view_action_run_get_data_3 = '    Argumentos de la ruta: {path}\n'
view_source_view_action_run_get_data_4 = '    Ingresar {arg}: '
view_source_view_action_run_get_data_5 = '    Argumentos del comando\n'

view_source_view_action_run_get_data_6 = 'No se ha encontrado el comando'

view_source_logs_export = 'Exportando logs en HOME.'
