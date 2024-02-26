import json

import requests

from examples.mockingbird.python.api_setup import api_v1_path, get_headers


def start_or_schedule_wf():
    print("Scheduled workflow exec")


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


wf_model_task_template = {
    "taskStrID": "0",
    "taskID": 0,
    "model": "gpt-4-0125-preview",
    "taskType": "analysis",
    "temperature": 1.0,
    "marginBuffer": 0.5,
    "taskGroup": "group",
    "taskName": "name",
    "maxTokens": 0,
    "tokenOverflowStrategy": "deduce",
    "prompt": "",
    "cycleCount": 1,
    "evalFns": []
}

if __name__ == '__main__':
    with open('templates/workflow.json', 'r') as file:
        data = json.load(file)

    data['workflowName'] = 'demo-workflow'
    data['workflowGroup'] = 'demo'

    data['stepSize'] = 5
    data['stepSizeUnit'] = 'minutes'

    pretty_data = json.dumps(data, indent=4)
    print(pretty_data)
    # create_wf(data)
