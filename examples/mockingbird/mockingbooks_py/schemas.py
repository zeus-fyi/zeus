import json  # Import the json module

import requests

from examples.mockingbird.mockingbooks_py.api_setup import get_headers, api_v1_path


def get_schemas():
    url = api_v1_path + "/schemas/ai"
    headers = get_headers()
    response = requests.get(url, headers=headers)
    if response.status_code == 200:
        pretty_data = json.dumps(response.json(), indent=4)
        print(pretty_data)
    else:
        print("Status Code:", response.status_code)
    return response.json()


def create_or_update_schema(schema):
    url = api_v1_path + "/schemas/ai"
    headers = get_headers()
    response = requests.post(url, json=schema, headers=headers)
    if response.status_code == 200:
        print("Schema created successfully!")
        pretty_data = json.dumps(response.json(), indent=4)
        print(pretty_data)
    else:
        print(response.json())
        print("Failed to create schema. Status Code:", response.status_code)
    return response.json()


field_data_template = {
    "fieldID": 0,
    "fieldStrID": "0",
    "fieldName": "msg_id",
    "fieldDescription": "The analyzed id value.",
    "dataType": "string"
}

if __name__ == '__main__':
    with open('templates/schema.json', 'r') as file:
        data = json.load(file)

    data['schemaName'] = 'demo-schema'
    data['schemaGroup'] = 'demo'
    data['schemaDescription'] = 'A schema for the demo task.'
    data['isObjArray'] = True

    field_data_template['fieldName'] = 'msg_id'
    field_data_template['fieldDescription'] = 'The analyzed id value.'
    field_data_template['dataType'] = 'string'
    field_data_template['description'] = 'string'
    data['fields'] = [field_data_template]

    create_or_update_schema(data)
