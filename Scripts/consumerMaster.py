import psycopg2
import boto3

conn = psycopg2.connect(
    host="prod-1.csz2smpfztf7.us-east-1.rds.amazonaws.com",
    database="maindb",
    user="registerio",
    password="registera")

cur = conn.cursor()
cur.execute("SELECT index from \"sqs queues\" WHERE index LIKE '07%'")

rows = cur.fetchall()
print(len(rows))
client = boto3.client('lambda')
for i, row in enumerate(rows):
    index = row[0]
    response = client.create_function(
        FunctionName="Consumer-{}".format(index),
        Role='arn:aws:iam::952897923483:role/Consumer',
        Code={
            'ImageUri': '952897923483.dkr.ecr.us-east-1.amazonaws.com/class-validation-consumer-lambda@sha256:85aeabddb9181200aa7b4a679b3c27eca9e0980374c70f3f3d0bd2dbc4a8dfc4'
        },
        Description="Consumer node for class index {}".format(index),
        Timeout=3,
        MemorySize=128,
        Publish=True,
        VpcConfig={
            'SubnetIds': [
                'subnet-02b72d28c0371cbad',
            ],
            'SecurityGroupIds': [
                'sg-0cd490b93d401aa72',
            ]
        },
        PackageType='Image',
        Tags={
            'INDEX': index
        },
        ImageConfig={
            'Command': [
                index,
            ]
        },
    )
    response = client.put_function_concurrency(
        FunctionName="Consumer-{}".format(index),
        ReservedConcurrentExecutions=1
    )