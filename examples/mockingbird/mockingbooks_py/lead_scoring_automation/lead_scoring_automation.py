import json

from examples.mockingbird.mockingbooks_py.agg_tasks import create_agg_task
from examples.mockingbird.mockingbooks_py.analysis_tasks import create_analysis_task
from examples.mockingbird.mockingbooks_py.evals import create_or_update_eval
from examples.mockingbird.mockingbooks_py.schemas import create_or_update_schema
from examples.mockingbird.mockingbooks_py.triggers import create_or_update_trigger
from examples.mockingbird.mockingbooks_py.workflows import create_wf
from examples.mockingbird.mockingbooks_py.workflows_examples import wf_model_task_template


def create_llm_workflows_schema():
    # Create a schema for LLM workflow lead scoring
    with open('lead_scoring_templates/llm_workflow_automation.txt', 'r') as file:
        llm_workflow_automation_lead_scoring = file.read()

    with open('mocks/schema_lead_score.json', 'r') as file:
        llm_workflow_automation_score_json = json.load(file)
        llm_workflow_automation_score_json['schemaName'] = 'lead_scoring_llm_workflows'
        llm_workflow_automation_score_json['fields'] = [
            {
                "fieldID": 0,
                "fieldStrID": "0",
                "fieldName": "lead_score",
                "fieldDescription": llm_workflow_automation_lead_scoring,
                "dataType": "integer"
            }
        ]
    return create_or_update_schema(llm_workflow_automation_score_json)


def create_developer_platform_schema():
    # Create a schema for developer platform workflow lead scoring
    with open('lead_scoring_templates/developer_platform.txt', 'r') as file:
        developer_platform_lead_scoring = file.read()

    with open('mocks/schema_lead_score.json', 'r') as file:
        developer_platform_lead__score_json = json.load(file)
        developer_platform_lead__score_json['fields'] = [
            {
                "fieldID": 0,
                "fieldStrID": "0",
                "fieldName": "lead_score",
                "fieldDescription": developer_platform_lead_scoring,
                "dataType": "integer"
            }
        ]
    return create_or_update_schema(developer_platform_lead__score_json)


def create_ef_filter_on_score():
    with open('mocks/eval_fn_trigger.json', 'r') as file:
        eval_fn_pl = json.load(file)
    create_or_update_eval(eval_fn_pl)


def create_trigger():
    with open('mocks/trigger_post_entities_data.json', 'r') as file:
        tlt = json.load(file)
    create_or_update_trigger(tlt)


def create_like_tweets_wf():
    with open('mocks/workflow.json', 'r') as file:
        wf_json = json.load(file)
    create_wf(wf_json)


def create_llm_wf_scoring_analysis_task(schema):
    with open('mocks/analysis.json', 'r') as file:
        at = json.load(file)
    at['schemas'] = [schema]
    at['taskName'] = 'llm_wf_scoring'
    create_analysis_task(at)


def create_dev_platform_scoring_analysis_task(schema):
    with open('mocks/analysis.json', 'r') as file:
        at = json.load(file)
    at['schemas'] = [schema]
    at['taskName'] = 'developer_platform_scoring'
    create_analysis_task(at)


def create_agg_msg_dev_platform():
    with open('mocks/aggregation.json', 'r') as file:
        at = json.load(file)
    with open('icp_templates/developer_platform.txt', 'r') as file:
        dev_platform_lead_scoring_icp_info = file.read()

    at['taskName'] = 'llm_wf_agg_msg_suggestion'
    at['prompt'] = dev_platform_lead_scoring_icp_info
    create_agg_task(at)


def create_agg_msg_llm_wfs():
    with open('mocks/aggregation.json', 'r') as file:
        at = json.load(file)
    with open('icp_templates/llm_workflow_automation.txt', 'r') as file:
        dev_platform_lead_scoring_icp_info = file.read()

    at['taskName'] = 'llm_wf_agg_msg_suggestion'
    at['prompt'] = dev_platform_lead_scoring_icp_info
    # create_agg_task(at)


def create_lead_score_wf():
    with open('mocks/workflow.json', 'r') as file:
        jdata = json.load(file)

    # Add a task to the workflow
    task_str_id = '0'
    wf_model_task_template['taskStrID'] = task_str_id
    jdata['models'][task_str_id] = wf_model_task_template

    # Get eval fn template
    with open('mocks/eval_fn_trigger.json', 'r') as file:
        ef_data = json.load(file)

    # Add an eval to the workflow
    jdata['evalsMap'] = {
        ef_data['evalStrID']: ef_data
    }
    jdata['evalTasksMap'] = {
        task_str_id: {
            ef_data['evalStrID']: True
        }
    }
    # Add an aggregate task to the workflow
    agg_task_str_id = '0'
    wf_model_task_template['taskType'] = 'aggregate'
    wf_model_task_template['taskStrID'] = agg_task_str_id
    jdata['models'][agg_task_str_id] = wf_model_task_template

    jdata['aggregateSubTasksMap'] = {
        agg_task_str_id: {
            task_str_id: True
        }
    }
    pretty_data = json.dumps(jdata, indent=4)
    print(pretty_data)
    create_wf(jdata)


def create_scoring_wfs():
    # create llm schema def
    llm_schema = create_llm_workflows_schema()
    # create llm scoring analysis task
    create_llm_wf_scoring_analysis_task(llm_schema)

    # create dev schema def
    dev_schema = create_developer_platform_schema()
    # create dev scoring analysis task
    create_dev_platform_scoring_analysis_task(dev_schema)

    # create a aggregation task llm
    create_agg_msg_llm_wfs()

    # create a aggregation task dev platform
    create_agg_msg_dev_platform()

    # create a trigger
    create_trigger()

    # create eval fn that triggers a tweet like POST request
    create_ef_filter_on_score()

    # create a workflow that uses the eval fn
    create_lead_score_wf()


if __name__ == '__main__':
    pass
    # search_entities_f = EntitiesFilter()
    #
    # pretty_data1 = search_entities(search_entities_f)
    # pretty_data2 = json.dumps(pretty_data1, indent=4)
    # print(pretty_data2)
    #
    # for v in pretty_data1:
    #     print(v)
    #     print('---')
