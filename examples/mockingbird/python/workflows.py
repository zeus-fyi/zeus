import json

import requests

from examples.mockingbird.python.api_setup import api_v1_path, get_headers


def start_or_schedule_wf(wf_exec_params):
    url = api_v1_path + "/workflows/ai/actions"
    headers = get_headers()
    response = requests.post(url, json=wf_exec_params, headers=headers)
    # Check the response status
    if response.status_code == 200:
        print("Workflow created successfully!")
    else:
        print(response.json())
        print("Failed to create workflow. Status Code:", response.status_code)


def create_wf(wf):
    url = api_v1_path + "/workflows/ai"
    headers = get_headers()
    response = requests.post(url, json=wf, headers=headers)
    # Check the response status
    if response.status_code == 200:
        print("Workflow created successfully!")
    else:
        print(response.json())
        print("Failed to create workflow. Status Code:", response.status_code)


if __name__ == '__main__':
    with open('templates/exec_wf.json', 'r') as file:
        payload = json.load(file)

    payload['workflows'] = {
        "workflowName": "wf-name-example",
        "workflowGroup": "demo"
    }
    start_or_schedule_wf(payload)
