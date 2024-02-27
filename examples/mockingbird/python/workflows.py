import json

import requests

from examples.mockingbird.python.api_setup import api_v1_path, get_headers


def start_or_schedule_wf(wf_exec_params):
    url = api_v1_path + "/workflows/ai/actions"
    headers = get_headers()
    response = requests.post(url, json=wf_exec_params, headers=headers)
    # Check the response status
    if response.status_code == 200:
        print(response.json())
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
        print(response.json())
        print("Workflow created successfully!")
    else:
        print(response.json())
        print("Failed to create workflow. Status Code:", response.status_code)


wf_exec_template = {
    "action": "start",
    "unixStartTime": 0,
    "duration": 1,
    "durationUnit": "cycles",
    "customBasePeriod": True,
    "customBasePeriodStepSize": 30,
    "customBasePeriodStepSizeUnit": "minutes",
    "workflows": []
}

wf_item_details = {
    "workflowName": "",
}

if __name__ == '__main__':
    # Starts a workflow
    wf_item_details['workflowName'] = 'demo-analysis-only-workflow'
    wf_exec_template['workflows'] = [wf_item_details]

    pretty_data = json.dumps(wf_exec_template, indent=4)
    print(pretty_data)
    start_or_schedule_wf(wf_exec_template)
