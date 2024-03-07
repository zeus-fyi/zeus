import json
import os

from examples.mockingbird.mockingbooks_py.agg_tasks import create_agg_task
from examples.mockingbird.mockingbooks_py.analysis_tasks import create_analysis_task
from examples.mockingbird.mockingbooks_py.evals import create_or_update_eval
from examples.mockingbird.mockingbooks_py.google_search_regex.dynamic_google_search import poll_run
from examples.mockingbird.mockingbooks_py.schemas import create_or_update_schema
from examples.mockingbird.mockingbooks_py.workflows import create_wf, wf_exec_template, start_or_schedule_wf
from examples.mockingbird.mockingbooks_py.workflows_examples import wf_model_task_template, wf_model_agg_task_template


def create_llm_workflows_schema():
    # Create a schema for LLM workflow lead scoring
    with open('lead_scoring_templates/llm_workflows_automation.txt', 'r') as file:
        llm_workflow_automation_lead_scoring = file.read()

    with open('mocks/llm_wf_schema_lead_score.json', 'r') as file:
        llm_workflow_automation_score_json = json.load(file)
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

    with open('mocks/dev_platform_schema_lead_score.json', 'r') as file:
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


def create_ef_filter_on_score_dev_platform():
    with open('mocks/eval_fn_dev_platform.json', 'r') as file:
        eval_fn_pl = json.load(file)
    create_or_update_eval(eval_fn_pl)


def create_ef_filter_on_score_llm_wfs():
    with open('mocks/eval_fn_llm_wfs.json', 'r') as file:
        eval_fn_pl = json.load(file)
    create_or_update_eval(eval_fn_pl)


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


def create_agg_msg_llm_wfs_platform():
    with open('mocks/agg_llm_wfs.json', 'r') as file:
        at = json.load(file)
    with open('icp_templates/mockingbird.txt', 'r') as file:
        llm_workflows_automation_icp_info = file.read()
    return create_agg(at, llm_workflows_automation_icp_info)


def create_agg_msg_dev_platform():
    with open('mocks/agg_dev_platform.json', 'r') as file:
        at = json.load(file)
    with open('icp_templates/developer_platform.txt', 'r') as file:
        dev_platform_lead_scoring_icp_info = file.read()
    return create_agg(at, dev_platform_lead_scoring_icp_info)


def create_agg(at, body_context):
    at['prompt'] = (
            "As a top-tier SaaS sales professional with a proven track record of swiftly closing deals, you understand the "
            "importance of first impressions. You've just received a promising lead, and you have our company product info and customer profile"
            "included to help assess from the context you're reviewing and need to craft an impactful cold outbound linkedIn message. Your approach combines deep insights into the potential client’s challenges with a concise "
            "showcase of how your software provides the perfect solution, all while creating a sense of urgency and "
            "exclusivity. Your goal is to not just introduce the software, but to create an immediate connection and "
            "schedule a follow-up demo or meeting, leveraging your legendary closing skills to make this opportunity "
            "impossible for the prospect to pass up. \n" + body_context)
    return create_agg_task(at)


def create_lead_score_llm_wf(task_str_id, agg_task_str_id):
    with open('mocks/workflow.json', 'r') as file:
        jdata = json.load(file)
    jdata['workflowName'] = 'llm_lead_scoring_wf'
    # Add a task to the workflow
    wf_model_task_template['taskStrID'] = task_str_id
    jdata['models'][task_str_id] = wf_model_task_template

    # Get eval fn template
    with open('mocks/eval_fn_llm_wfs.json', 'r') as file:
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
    wf_model_agg_task_template['taskStrID'] = agg_task_str_id
    jdata['models'][agg_task_str_id] = wf_model_agg_task_template

    jdata['aggregateSubTasksMap'] = {
        agg_task_str_id: {
            task_str_id: True
        }
    }
    pretty_data_out = json.dumps(jdata, indent=4)
    print(pretty_data_out)
    create_wf(jdata)


def create_lead_score_dp_wf(task_str_id, agg_task_str_id):
    with open('mocks/workflow.json', 'r') as file:
        jdata = json.load(file)

    jdata['workflowName'] = 'dp_lead_scoring_wf'
    # Add a task to the workflow
    wf_model_task_template['taskStrID'] = task_str_id
    jdata['models'][task_str_id] = wf_model_task_template

    # Get eval fn template
    with open('mocks/eval_fn_dev_platform.json', 'r') as file:
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
    wf_model_agg_task_template['taskStrID'] = agg_task_str_id
    jdata['models'][agg_task_str_id] = wf_model_agg_task_template

    jdata['aggregateSubTasksMap'] = {
        agg_task_str_id: {
            task_str_id: True
        }
    }
    pretty_data_out = json.dumps(jdata, indent=4)
    print(pretty_data_out)
    create_wf(jdata)


def create_scoring_wfs():
    # create llm schema def
    llm_schema_wf = create_llm_workflows_schema()
    # create llm scoring analysis task
    create_llm_wf_scoring_analysis_task(llm_schema_wf)

    # create dev schema def
    dev_schema_wf = create_developer_platform_schema()
    # create dev scoring analysis task
    create_dev_platform_scoring_analysis_task(dev_schema_wf)

    create_ef_filter_on_score_llm_wfs()
    create_ef_filter_on_score_dev_platform()

    # create an aggregation task for both
    create_agg_msg_llm_wfs_platform()
    create_agg_msg_dev_platform()

    # create eval fn that triggers a tweet like POST request
    # create a workflow that uses the eval fn
    llm_task_str_id = '1709526913053514000'
    llm_agg_task_str_id = '1709526980909318000'
    create_lead_score_llm_wf(llm_task_str_id, llm_agg_task_str_id)

    dev_task_str_id = '1709526931827736000'
    dev_agg_task_str_id = '1709526981067542000'
    create_lead_score_dp_wf(dev_task_str_id, dev_agg_task_str_id)


def run_dp_scoring_wf(entity, dry_run=False):
    # dp_lead_scoring_wf
    wf_item_details['workflowName'] = 'dp_lead_scoring_wf'
    wf_exec_template['workflows'] = [wf_item_details]
    tmp = {'developer_platform_scoring': {'replacePrompt': entity}}
    wf_exec_template['taskOverrides'] = tmp
    pretty_data = json.dumps(wf_exec_template, indent=4)
    print(pretty_data)

    if not dry_run:
        start_or_schedule_wf(wf_exec_template)


def run_llm_wfs_scoring_wf(entity, dry_run=False):
    # llm_lead_scoring_wfs
    wf_item_details['workflowName'] = 'llm_lead_scoring_wf'
    wf_exec_template['workflows'] = [wf_item_details]

    tmp = {'llm_wf_scoring': {'replacePrompt': entity}}
    wf_exec_template['taskOverrides'] = tmp

    pretty_data = json.dumps(wf_exec_template, indent=4)
    print(pretty_data)
    if not dry_run:
        start_or_schedule_wf(wf_exec_template)


# Function to categorize lead score
def categorize_score(score):
    if 23 <= score <= 28:
        return 'high'
    elif 15 <= score <= 22:
        return 'med'
    elif 8 <= score <= 14:
        return 'low'
    else:  # 0-7
        return 'Not a priority'


wf_item_details = {
    "workflowName": "",
}

if __name__ == '__main__':
    # iterate_on_matches()
    #
    # search_entities_f = EntitiesFilter(
    #     sinceUnixTimestamp=-80000,
    #     platform='linkedIn'
    # )
    #
    # entities_saved = search_entities(search_entities_f)
    # print(len(entities_saved))
    # pretty_data2 = json.dumps(entities_saved, indent=4)

    # for v in entities_saved:
    #     print(v)
    #     print('---')
    #
    # # for when you want to analyze a targeted entity/platform,
    # # 1. quick local find + target wf
    # # 2. local prototyping of the wf using real entity data without larger scale AI search
    #
    # dry_run_wf = False
    # for i, tgt in enumerate(entities_saved):
    #     tgt_entity = json.dumps(tgt)
    #     run_dp_scoring_wf(tgt_entity, dry_run_wf)
    #     run_llm_wfs_scoring_wf(tgt_entity, dry_run_wf)
    #     time.sleep(60)

    base_path = 'tmp'
    file_path = f'tmp/tmp2.txt'
    orchestration_results = {}
    empty_count = 0
    complete_count = 0
    # Open the file and read its contents
    # Read the file and process lines
    with open(file_path, 'r') as file:
        for line in file:
            clean_line = line.strip()
            res = poll_run(clean_line)

            if not res:
                empty_count += 1
                continue

            for item in res:
                if item and 'aggregatedEvalResults' in item and isinstance(item['aggregatedEvalResults'], list) and len(
                        item['aggregatedEvalResults']) > 0:
                    orchestration_name = item.get('orchestration', {}).get('type', 'Unknown')
                    # Path for the orchestration
                    orchestration_path = os.path.join(base_path, orchestration_name)

                    if not os.path.exists(orchestration_path):
                        os.makedirs(orchestration_path)

                    for evr in item['aggregatedEvalResults']:
                        if evr.get('evalMetricResult') and evr['evalMetricResult'].get('evalMetadata') and 'intValue' in \
                                evr['evalMetricResult']['evalMetadata']:
                            lead_score = evr['evalMetricResult']['evalMetadata']['intValue']
                            category = categorize_score(lead_score)
                            category_path = os.path.join(orchestration_path, category)

                            if not os.path.exists(category_path):
                                os.makedirs(category_path)

                            # Here you would write your file or data to the directory
                            # For example:
                            filename = f'{line}.json'
                            with open(os.path.join(category_path, filename), 'w') as result_file:
                                pretty_data = json.dumps(res, indent=4)
                                result_file.write(pretty_data)
                            complete_count += 1
                        else:
                            empty_count += 1
                else:
                    empty_count += 1

    # Print results
    for orchestration_name, categories in orchestration_results.items():
        print(f"Orchestration Name: {orchestration_name}")
        for category, count in categories.items():
            print(f"  {category}: {count}")
    print(f'Empty: {empty_count}, Complete: {complete_count}')
