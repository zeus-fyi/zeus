import json

from examples.mockingbird.mockingbooks_py.retrievals import create_or_update_retrieval
from examples.mockingbird.mockingbooks_py.schemas import create_or_update_schema
from examples.mockingbird.mockingbooks_py.workflows import create_wf


def schema_create():
    with open('mocks/schema.json', 'r') as file:
        data = json.load(file)
    create_or_update_schema(data)


def retrieval_create():
    with open('mocks/retrieval.json', 'r') as file:
        data = json.load(file)
    create_or_update_retrieval(data)


def create_google_regex_search_index_entities_wf(task_str_id, eval_str_id, agg_task_str_id, agg_eval_str_id=None):
    with open('mocks/workflow.json', 'r') as file:
        jdata = json.load(file)

    jdata['stepSize'] = 30
    jdata['stepSizeUnit'] = 'minutes'

    # Add a task to the workflow
    with open('mocks/analysis_task.json', 'r') as file:
        wf_analysis_task = json.load(file)

    wf_analysis_task['taskStrID'] = task_str_id
    wf_analysis_task['taskType'] = 'analysis'
    wf_analysis_task['modelType'] = 'gpt-4-0125-preview'
    jdata['models'][task_str_id] = wf_analysis_task

    pretty_data = json.dumps(jdata, indent=4)
    print(pretty_data)
    create_wf(jdata)


if __name__ == '__main__':
    # retrieval_create()
    # schema_create()
    pass
