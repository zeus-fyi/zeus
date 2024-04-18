import json

from examples.mockingbird.mockingbooks_py.analysis_tasks import create_analysis_task, get_task_id_by_name
from examples.mockingbird.mockingbooks_py.retrievals import create_or_update_retrieval, get_retrieval_id_by_name
from examples.mockingbird.mockingbooks_py.schemas import create_or_update_schema, get_schema_by_name
from examples.mockingbird.mockingbooks_py.workflows import create_wf


def schema_create():
    with open('mocks/schema.json', 'r') as file:
        data = json.load(file)
    create_or_update_schema(data)


def retrieval_create():
    with open('mocks/retrieval.json', 'r') as file:
        data = json.load(file)
    create_or_update_retrieval(data)


def analysis_create():
    with open('mocks/analysis.json', 'r') as file:
        data = json.load(file)
    data['schemas'] = [get_schema_by_name('website-analysis')]
    create_analysis_task(data)


def create_website_analysis_wf(retrieval_str_id, task_str_id):
    with open('mocks/workflow.json', 'r') as file:
        jdata = json.load(file)

    jdata['stepSize'] = 4
    jdata['stepSizeUnit'] = 'hours'

    # Add a task to the workflow
    with open('mocks/analysis.json', 'r') as file:
        wf_analysis_task = json.load(file)

    wf_analysis_task['taskStrID'] = task_str_id
    wf_analysis_task['taskType'] = 'analysis'
    # wf_analysis_task['modelType'] = 'gpt-4-0125-preview'
    jdata['models'][task_str_id] = wf_analysis_task
    jdata['analysisRetrievalsMap'] = {
        task_str_id: {
            retrieval_str_id: True,
        }
    }
    pretty_data = json.dumps(jdata, indent=4)
    print(pretty_data)
    create_wf(jdata)


if __name__ == '__main__':
    # retrieval_create()
    # schema_create()
    # analysis_create()
    ret_id = get_retrieval_id_by_name('website-analysis')
    task_id = get_task_id_by_name('website-analysis')
    create_website_analysis_wf(ret_id, task_id)
