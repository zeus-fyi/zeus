import json  # Import the json module

import requests

from examples.mockingbird.mockingbooks_py.api_setup import get_headers, api_v1_path
from examples.mockingbird.mockingbooks_py.triggers import trigger_function_template


def get_evals():
    url = api_v1_path + "/evals/ai"
    headers = get_headers()
    response = requests.get(url, headers=headers)
    if response.status_code == 200:
        data = response.json()
        pretty_data = json.dumps(data, indent=4)
        print(pretty_data)
    else:
        print("Status Code:", response.status_code)
    return response.json()


def create_or_update_eval(eval_fn):
    url = api_v1_path + "/evals/ai"
    headers = get_headers()
    response = requests.post(url, json=eval_fn, headers=headers)
    if response.status_code == 200:
        data = response.json()
        pretty_data = json.dumps(data, indent=4)
        print(pretty_data)
    else:
        print("Status Code:", response.status_code)
    return response.json()


if __name__ == '__main__':
    with open('templates/eval_fn.json', 'r') as file:
        eval_fn_pl = json.load(file)

    # for updating an existing eval function use id
    # eval_fn_str_id = '1709068762300906000'
    # eval_fn_pl['evalStrID'] = eval_fn_str_id

    eval_fn_pl['evalName'] = 'demo-eval-fn'
    eval_fn_pl['evalGroup'] = 'demo'
    eval_fn_pl['evalFormat'] = 'json'
    eval_fn_pl['evalModel'] = 'gpt-4-0125-preview'

    # Add schema
    with open('google_search_regex/schema.json', 'r') as file:
        efn_schema = json.load(file)

    efn_schema['schemaStrID'] = '1708624962876732000'
    efn_schema['fields'][0]['fieldStrID'] = '1708580020120821000'
    eval_fn_pl['schemas'] = [efn_schema]

    # Add trigger
    trigger_str_id = '1708624962876732000'
    trigger_function_template['triggerStrID'] = trigger_str_id
    eval_fn_pl['triggerFunctions'] = [trigger_function_template]

    pretty_task_data = json.dumps(eval_fn_pl, indent=4)
    print(pretty_task_data)

    create_or_update_eval(eval_fn_pl)
