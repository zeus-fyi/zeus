import json

from examples.mockingbird.mockingbooks_py.analysis_tasks import create_analysis_task
from examples.mockingbird.mockingbooks_py.retrievals import create_or_update_retrieval, get_retrieval_id_by_name
from examples.mockingbird.mockingbooks_py.schemas import create_or_update_schema, get_schema_by_name
from examples.mockingbird.mockingbooks_py.triggers import create_or_update_trigger


def retrieval_create():
    with open('mocks/retrieval.json', 'r') as file:
        data = json.load(file)
    create_or_update_retrieval(data)


def schema_create():
    with open('mocks/schema.json', 'r') as file:
        data = json.load(file)
    create_or_update_schema(data)


def analysis_create():
    with open('mocks/analysis_task.json', 'r') as file:
        data = json.load(file)
    data['schemas'] = [get_schema_by_name('linked-in-rapid-api-profile-qps')]
    create_analysis_task(data)


def trigger_create():
    with open('mocks/trigger.json', 'r') as file:
        data = json.load(file)
    for tr in data['triggerRetrievals']:
        if tr['retrievalName'] == 'linkedin-profile':
            tr['retrievalStrID'] = get_retrieval_id_by_name('linkedin-profile')
    create_or_update_trigger(data)


if __name__ == '__main__':
    # retrieval_create()
    # schema_create()
    # analysis_create()
    # trigger_create()
    pass
