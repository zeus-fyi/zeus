import json

from examples.mockingbird.mockingbooks_py.analysis_tasks import get_task_id_by_name, create_analysis_task
from examples.mockingbird.mockingbooks_py.retrievals import create_or_update_retrieval, get_retrieval_id_by_name
from examples.mockingbird.mockingbooks_py.schemas import create_or_update_schema
from examples.mockingbird.mockingbooks_py.workflows import create_wf


# currently unused
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
    create_analysis_task(data)


def create_email_validation_wf():
    with open('mocks/workflow.json', 'r') as file:
        jdata = json.load(file)

    jdata['stepSize'] = 30
    jdata['stepSizeUnit'] = 'minutes'

    # Add a task to the workflow
    with open('mocks/analysis.json', 'r') as file:
        wf_analysis_task = json.load(file)
    task_str_id = get_task_id_by_name(wf_analysis_task['taskName'])

    with open('mocks/retrieval.json', 'r') as file:
        data = json.load(file)
    retrieval_str_id = get_retrieval_id_by_name(data['retrievalName'])
    wf_analysis_task['taskStrID'] = task_str_id
    wf_analysis_task['taskType'] = 'analysis'
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
    # analysis_create()
    create_email_validation_wf()
    pass
