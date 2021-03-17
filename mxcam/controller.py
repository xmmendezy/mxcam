# -*- coding: utf-8 -*-

from sys import exit
from pathlib import Path
from os import system, name
from signal import signal, SIGINT
from time import sleep
from json import dump, load, loads
from terminaltables import AsciiTable
from re import findall
from datetime import datetime
from requests import post
from .view import (
    error_help,
    separator_help,
    back_table_helper,
    cancel_input_helper,
    log_default,
    view_main,
    view_user, view_user_create_0, view_user_create_1, view_user_create_1_1, view_user_create_1_2, view_user_create_1_3,
    view_user_update_0, view_user_update_0_1, view_user_update_0_2, view_user_update_0_3, view_user_update_1,
    view_command, view_command_run_0, view_command_create_0, view_command_create_0_1, view_command_create_0_2, view_command_create_0_3,
    view_command_create_0_4, view_command_create_0_5, view_command_create_1, view_command_create_2,
    view_command_list_0, view_command_list_1, view_command_delete_0, view_command_delete_1, view_command_delete_2, view_command_delete_3,
    view_source, view_source_create_0, view_source_create_0_1, view_source_create_0_2, view_source_create_1, view_source_list_0, view_source_list_1,
    view_source_delete_0, view_source_delete_1, view_source_delete_2, view_source_delete_3,
    view_source_view_0, view_source_view_1, view_source_view_2, view_source_view_action,
    view_source_view_action_add_0, view_source_view_action_add_0_1, view_source_view_action_add_0_2, view_source_view_action_add_1,
    view_source_view_action_list_0, view_source_view_action_list_1, view_source_view_action_delete_0, view_source_view_action_delete_1,
    view_source_view_action_delete_2, view_source_view_action_delete_3, view_source_source_export_0, view_source_source_export_1,
    view_source_source_export_2, view_source_source_export_3, view_source_source_import_0, view_source_source_import_1,
    view_source_source_import_2, view_source_source_import_3, view_source_source_import_4,
    view_source_view_action_run_0, view_source_view_action_run_1, view_source_view_action_run_2, view_source_view_action_run_3,
    view_source_view_action_run_get_data_1, view_source_view_action_run_get_data_2, view_source_view_action_run_get_data_3,
    view_source_view_action_run_get_data_4, view_source_view_action_run_get_data_5, view_source_view_action_run_get_data_6,
    view_source_logs_export,
)
from .model import (
    User, Command, Source, Ip, Log,
)

__all__ = [
    'controller_main',
]

# Utils

source_select: Source = None
home = str(Path.home())


def handler_default(signal_received, frame):
    pass


signal(SIGINT, handler_default)


def clear():
    if name == 'nt':
        _ = system('cls')
    else:
        _ = system('clear')


def loading(message='Cargando', times=2):
    for l in '/-\\|' * times:
        clear()
        print(f'\n    {message}   {l}')
        sleep(0.25)


def printProgressBar(iteration, total, prefix='', suffix='', decimals=1, length=100, fill='█', logs=[]):
    percent = ("{0:." + str(decimals) + "f}").format(100 * (iteration / float(total)))
    filledLength = int(length * iteration // total)
    bar = fill * filledLength + '-' * (length - filledLength)
    clear()
    print('    %s |%s| %s%% %s' % (prefix, bar, percent, suffix), end='\n\n')
    for log in logs:
        print(f'    - {log}')
    if iteration == total:
        print()


def controller_base(view=view_main, actions={}):
    func = None
    actions = dict(actions, **{
        'exit': controller_exit,
        'e': controller_exit,
    })
    while True:
        print(view)
        action = input('    Acción: ')
        func = actions.get(action.lower(), None)
        if func:
            break
        else:
            print(error_help)
    func()


# Controllers

def controller_exit():
    exit()


def controller_main():
    clear()
    controller_base(view_main, {
        'user': controller_user,
        'u': controller_user,
        'command': controller_command,
        'c': controller_command,
        'source': controller_source,
        's': controller_source,
        'logs': controller_logs_export,
        'l': controller_logs_export,
    })


def controller_user(do_clear=True):
    if do_clear:
        clear()
    controller_base(view_user, {
        'create': controller_user_create,
        'c': controller_user_create,
        'update': controller_user_update,
        'u': controller_user_update,
        'back': controller_main,
        'b': controller_main,
    })


def controller_user_create(do_clear=True):
    if do_clear:
        clear()
    if User.query.all():
        loading(view_user_create_0)
    else:
        while True:
            print(view_user_create_1)
            username = input(view_user_create_1_1)
            if username.lower() == 'cancel':
                return controller_user()
            password = input(view_user_create_1_2)
            if password.lower() == 'cancel':
                return controller_user()
            if username and password:
                break
        user = User(username, password)
        loading(view_user_create_1_3)
        user.save()
    controller_user()


def controller_user_update(do_clear=True):
    if do_clear:
        clear()
    if User.query.all():
        print(view_user_update_0)
        username = input(view_user_update_0_1)
        if username.lower() == 'cancel':
            return controller_user()
        password = input(view_user_update_0_2)
        if password.lower() == 'cancel':
            return controller_user()
        if username or password:
            user: User = User.query.first()
            if username:
                user.username = username
            if password:
                user.password = password
            user.save()
            loading(view_user_update_0_3, 1)
    else:
        loading(view_user_update_1)
    controller_user()


def controller_command(do_clear=True):
    if do_clear:
        clear()
    controller_base(view_command, {
        'create': controller_command_create,
        'c': controller_command_create,
        'list': controller_command_list,
        'l': controller_command_list,
        'delete': controller_command_delete,
        'd': controller_command_delete,
        'back': controller_main,
        'b': controller_main,
    })


def controller_command_create(do_clear=True):
    if do_clear:
        clear()
    print(view_command_create_0)
    name = input(view_command_create_0_1)
    if name.lower() == 'cancel':
        return controller_command()
    path = input(view_command_create_0_2)
    if path.lower() == 'cancel':
        return controller_command()
    if not path.startswith('/'):
        path = f'/{path}'
    print(view_command_create_0_3)
    value = ''
    while True:
        line = input()
        if line.lower() == 'cancel':
            return controller_command()
        if line:
            if value:
                value = f'{value}\n{line}'
            else:
                value = line
        else:
            break
    print(view_command_create_0_4)
    message_logs = ''
    while True:
        line = input()
        if line.lower() == 'cancel':
            return controller_command()
        if line:
            if message_logs:
                message_logs = f'{message_logs}\n{line}'
            else:
                message_logs = line
        else:
            break
    default_message = input(view_command_create_0_5)
    if default_message.lower() == 'cancel':
        return controller_command()
    if name and path and value:
        if not message_logs.startswith('{'):
            message_logs = f'{{ {message_logs}'
        if message_logs.endswith(','):
            message_logs = message_logs[:-1]
        if message_logs.endswith(',}'):
            message_logs = f'{message_logs[:-2]} }}'
        if not message_logs.endswith('}'):
            message_logs = f'{message_logs} }}'
        message_logs = message_logs.replace('\'', '"')
        comand = Command(name, path, value, message_logs, default_message)
        comand.save()
        loading(view_command_create_1)
    else:
        loading(view_command_create_2)
    controller_command()


def controller_command_list(do_clear=True):
    if do_clear:
        clear()
    commands = Command.query.all()
    if commands:
        print(view_command_list_0)
        table = AsciiTable([['id', 'Nombre']] + [[command.id, command.name] for command in commands])
        table.padding_left = 2
        table.padding_right = 2
        print('\n'.join([f'\t{line}' for line in table.table.split('\n')]))
    else:
        print(view_command_list_1)
    input(back_table_helper)
    controller_command()


def controller_command_delete(do_clear=True):
    if do_clear:
        clear()
    print(view_command_delete_0)
    search = input(view_command_delete_1)
    if search.lower() == 'cancel':
        return controller_command()
    command = Command.query.filter_by(id=search).first()
    if command:
        command.delete()
        loading(view_command_delete_2)
    else:
        loading(view_command_delete_3)
    controller_command()


def controller_source(do_clear=True):
    if do_clear:
        clear()
    global source_select
    source_select = None
    controller_base(view_source, {
        'create': controller_source_create,
        'c': controller_source_create,
        'list': controller_source_list,
        'l': controller_source_list,
        'delete': controller_source_delete,
        'd': controller_source_delete,
        'view': controller_source_view,
        'v': controller_source_view,
        'import_file': controller_source_import,
        'if': controller_source_import,
        'export_file': controller_source_export,
        'ef': controller_source_export,
        'back': controller_main,
        'b': controller_main,
    })


def controller_source_create(do_clear=True):
    if do_clear:
        clear()
    print(view_source_create_0)
    while True:
        name = input(view_source_create_0_1)
        if name.lower() == 'cancel':
            return controller_source()
        if name:
            break
    comment = input(view_source_create_0_2)
    if comment.lower() == 'cancel':
        return controller_source()
    source = Source(name, comment)
    source.save()
    loading(view_source_create_1)
    controller_source()


def controller_source_list(do_clear=True):
    if do_clear:
        clear()
    sources = Source.query.all()
    if sources:
        print(view_source_list_0)
        table = AsciiTable([['id', 'Nombre', 'Comentario']] + [[source.id, source.name, source.comment] for source in sources])
        table.padding_left = 2
        table.padding_right = 2
        print('\n'.join([f'\t{line}' for line in table.table.split('\n')]))
    else:
        print(view_source_list_1)
    input(back_table_helper)
    controller_source()


def controller_source_delete(do_clear=True):
    if do_clear:
        clear()
    print(view_source_delete_0)
    search = input(view_source_delete_1)
    if search.lower() == 'cancel':
        return controller_source()
    source = Source.query.filter_by(id=search).first()
    if source:
        source.delete()
        loading(view_source_delete_2)
    else:
        loading(view_source_delete_3)
    controller_source()


def controller_source_view(do_clear=True):
    if do_clear:
        clear()
    global source_select
    print(view_source_view_0)
    search = input(view_source_view_1)
    if search.lower() == 'cancel':
        return controller_source()
    source = Source.query.filter_by(id=search).first()
    if source:
        source_select = source
        controller_source_view_action()
    else:
        loading(view_source_view_2)
    controller_source()


def controller_source_view_action(do_clear=True):
    if do_clear:
        clear()
    global source_select
    controller_base(view_source_view_action.format(**{'source': source_select.name}), {
        'add': controller_source_view_action_add,
        'a': controller_source_view_action_add,
        'list': controller_source_view_action_list,
        'l': controller_source_view_action_list,
        'delete': controller_source_view_action_delete,
        'd': controller_source_view_action_delete,
        'run': controller_source_view_action_run,
        'r': controller_source_view_action_run,
        'run_ip': controller_source_view_action_run_ip,
        'ri': controller_source_view_action_run_ip,
        'export_file': controller_source_export,
        'ef': controller_source_export,
        'back': controller_source,
        'b': controller_source,
    })


def controller_source_view_action_add(do_clear=True):
    if do_clear:
        clear()
    global source_select
    print(view_source_view_action_add_0.format(**{'source': source_select.name}))
    while True:
        ip = input(view_source_view_action_add_0_1)
        if ip.lower() == 'cancel':
            return controller_source_view_action()
        port = input(view_source_view_action_add_0_2)
        if port.lower() == 'cancel':
            return controller_source_view_action()
        if ip and port:
            break
    ip_el = Ip(ip, port, source_select.id)
    ip_el.save()
    loading(view_source_view_action_add_1)
    controller_source_view_action()


def controller_source_view_action_list(do_clear=True):
    if do_clear:
        clear()
    global source_select
    ips = Ip.query.filter(Ip.source.has(id=source_select.id)).all()
    if ips:
        print(view_source_view_action_list_0.format(**{'source': source_select.name}))
        table = AsciiTable([['id', 'IP/Dirección', 'Puerto']] + [[ip.id, ip.ip, ip.port] for ip in ips])
        table.padding_left = 2
        table.padding_right = 2
        print('\n'.join([f'\t{line}' for line in table.table.split('\n')]))
    else:
        print(view_source_view_action_list_1.format(**{'source': source_select.name}))
    input(back_table_helper)
    controller_source_view_action()


def controller_source_view_action_delete(do_clear=True):
    if do_clear:
        clear()
    global source_select
    print(view_source_view_action_delete_0.format(**{'source': source_select.name}))
    search = input(view_source_view_action_delete_1)
    if search.lower() == 'cancel':
        return controller_source_view_action()
    ip = Ip.query.filter(Ip.source.has(id=source_select.id)).filter_by(id=search).first()
    if ip:
        ip.delete()
        loading(view_source_view_action_delete_2)
    else:
        loading(view_source_view_action_delete_3)
    controller_source_view_action()


def controller_source_view_action_run(do_clear=True):
    if do_clear:
        clear()
    global source_select
    user = User.query.first()
    if not user:
        loading(view_source_view_action_run_0)
        return controller_source_view_action()
    print(view_source_view_action_run_1.format(**{'source': source_select.name}))
    path, command, message_logs, default_message = get_data_run_source()
    if not path and not command:
        return controller_source_view_action()
    ips = Ip.query.filter(Ip.source.has(id=source_select.id)).all()
    exec_on_ips(ips, user, path, command, message_logs, default_message)
    controller_source_view_action(False)


def controller_source_view_action_run_ip(do_clear=True):
    if do_clear:
        clear()
    global source_select
    user = User.query.first()
    if not user:
        loading(view_source_view_action_run_0)
        return controller_source_view_action()
    print(view_source_view_action_run_1.format(**{'source': source_select.name}))
    search = input(view_source_view_action_run_2)
    if search.lower() == 'cancel':
        return controller_source_view_action()
    ip = Ip.query.filter(Ip.source.has(id=source_select.id)).filter_by(id=search).first()
    if not ip:
        loading(view_source_view_action_run_3)
        return controller_source_view_action()
    path, command, message_logs, default_message = get_data_run_source()
    if not path and not command:
        return controller_source_view_action()
    exec_on_ips([ip], user, path, command, message_logs, default_message)
    controller_source_view_action(False)


def get_data_run_source() -> [str, str, dict, str]:
    search = input(view_source_view_action_run_get_data_1)
    if search.lower() == 'cancel':
        return controller_source_view_action()
    command: Command = Command.query.filter_by(id=search).first()
    if command:
        print(view_source_view_action_run_get_data_2.format(**{'command': command.name}))
        args_path = [l[1:-1] for l in findall('\{.*?\}', command.path)]
        args_path_dict = {}
        if args_path:
            print(view_source_view_action_run_get_data_3.format(**{'path': command.path}))
            for arg in args_path:
                value = input(view_source_view_action_run_get_data_4.format(**{'arg': arg}))
                if value.lower() == 'cancel':
                    return controller_source_view_action()
                args_path_dict[arg] = value
        args_value = [l[1:-1] for l in findall('\{.*?\}', command.value)]
        args_value_dict = {}
        if args_value:
            print(view_source_view_action_run_get_data_5)
            for arg in args_value:
                value = input(view_source_view_action_run_get_data_4.format(**{'arg': arg}))
                if value.lower() == 'cancel':
                    return controller_source_view_action()
                args_value_dict[arg] = value
        default_message = command.default_message if command.default_message else log_default
        return [
            command.path.format(**args_path_dict),
            command.value.format(**args_value_dict),
            loads(command.message_logs),
            default_message,
        ]
    else:
        loading(view_source_view_action_run_get_data_6)
        return ['', '', {}, '']


def exec_run_source(user: User, ip: Ip, path: str, command: str):
    if ip.ip.startswith('http'):
        target_url = f'{ip.ip}:{ip.port}{path}'
    else:
        target_url = f'http://{ip.ip}:{ip.port}{path}'
    try:
        response = post(url=target_url, auth=(user.username, user.password), data=command, headers={'Content-Type': 'application/octet-stream'})
        return response.status_code
    except:
        return 500


def exec_on_ips(ips, user: User, path: str, command: str, message_logs: dict, default_message: str):
    clear()
    printProgressBar(0, len(ips), prefix='Progreso:', suffix='Completo', length=10 * len(ips))
    logs = []
    for i, ip in enumerate(ips):
        command_i = command
        arg_replace = findall('\{.*?\}', command)
        if arg_replace:
            arg_evals = [str(eval(s)) for s in [l.replace('i', str(i)) for l in [l[1:-1] for l in arg_replace]]]
            for r, s in zip(arg_replace, arg_evals):
                command_i = command_i.replace(r, s)
        status = exec_run_source(user, ip, path, command_i)
        log = Log(
            message_logs.get(str(status), default_message).format(**{
                'now': datetime.now().strftime('%Y-%m-%d %H:%M:%S'),
                'date': datetime.now().strftime('%Y-%m-%d'),
                'time': datetime.now().strftime('%H:%M:%S'),
                'ip': ip.ip,
                'port': ip.port,
                'status': status
            })
        )
        log.save()
        logs.append(log.value)
        printProgressBar(i + 1, len(ips), prefix='Progreso:', suffix='Completo', length=10 * len(ips), logs=logs)
    print(f'\n    {separator_help}\n')


def controller_source_import(do_clear=True):
    if do_clear:
        clear()
    print(view_source_source_import_0)
    search = input(view_source_source_import_1)
    if search.lower() == 'cancel':
        return controller_source()
    try:
        with open(search, 'r') as inputfile:
            source_ser = load(inputfile)
            if source_ser['name']:
                source = Source(source_ser['name'], source_ser.get('comment', ''))
                source.save()
                for ip in source_ser.get('ips', []):
                    ip_el = Ip(ip['ip'], ip['port'], source.id)
                    ip_el.save()
                loading(view_source_source_import_2)
            else:
                loading(view_source_source_import_3)
    except:
        loading(view_source_source_import_4)
    controller_source()


def controller_source_export(do_clear=True):
    if do_clear:
        clear()
    global source_select
    global home
    if source_select:
        to_action = True
        source = source_select
    else:
        to_action = False
        print(view_source_source_export_0)
        search = input(view_source_source_export_1)
        if search.lower() == 'cancel':
            return controller_source()
        source = Source.query.filter_by(id=search).first()
        if not source:
            loading(view_source_source_export_2)
            return controller_source()
    source_ser = {
        'name': source.name,
        'comment': source.comment,
        'ips': []
    }
    ips = Ip.query.filter(Ip.source.has(id=source.id)).all()
    for ip in ips:
        source_ser['ips'].append({
            'ip': ip.ip,
            'port': ip.port,
        })
    with open(f'{home}/{source.name}-{source.id}.json', 'w') as outfile:
        dump(source_ser, outfile, indent=4)
    loading(view_source_source_export_3)
    if to_action:
        controller_source_view_action()
    else:
        controller_source()


def controller_logs_export(do_clear=True):
    if do_clear:
        clear()
    global home
    logs = Log.query.all()
    now = datetime.now().strftime('%Y-%m-%d_%H:%M:%S')
    with open(f'{home}/logs-mxcam-{now}.txt', 'w') as outfile:
        for log in logs:
            outfile.write(f'{log.value}\n')
    loading(view_source_logs_export)
    controller_main()
