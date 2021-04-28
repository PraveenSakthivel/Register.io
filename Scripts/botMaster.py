import boto3

client = boto3.client('ecs')
for i in range (2):
    response = client.run_task(
        cluster='arn:aws:ecs:us-east-1:952897923483:cluster/cv-consumer-prod-1',
        count=1,
        enableExecuteCommand=False,
        launchType='FARGATE',
        networkConfiguration={
            'awsvpcConfiguration': {
                'subnets': [
                    'subnet-09cd406fc7c061218',
                    'subnet-04a907094dfab1bf0'
                ],
                'securityGroups': [
                    'sg-0cd490b93d401aa72',
                ],
                'assignPublicIp': 'ENABLED'
            }
        },
        overrides={
            'containerOverrides': [
                {
                    'name': 'clientBot',
                    'environment': [
                        {
                            'name': 'START',
                            'value': '{}'.format(i*200)
                        },
                    ]
                },
            ]
        },
        taskDefinition='bot:2'
    )
