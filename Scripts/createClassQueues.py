import psycopg2
import boto3

conn = psycopg2.connect(
    host="prod-1.csz2smpfztf7.us-east-1.rds.amazonaws.com",
    database="maindb",
    user="registerio",
    password="registera")

cur = conn.cursor()

cur.execute("SELECT index from soc WHERE department = 119")

row = cur.fetchone()

pairs = []

while row is not None:
    index = row[0]
    queue = index + ".fifo"
    client = boto3.client('sqs')
    response = client.create_queue(
        QueueName=queue,
        Attributes={
            'ReceiveMessageWaitTimeSeconds': '5',
            'VisibilityTimeout': '5',
            'FifoQueue': 'true',
            'DeduplicationScope': 'messageGroup',
            'FifoThroughputLimit': 'perMessageGroupId'
        },
    )
    url = response['QueueUrl']
    pairs.append((index,url))
    row = cur.fetchone()

for (index, url) in pairs:
    cur.execute("INSERT INTO \"sqs queues\"(index, url) VALUES (%s,%s);", (index,url))


conn.commit()

cur.close()
conn.close()

