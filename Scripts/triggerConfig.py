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
    response = client.create_event_source_mapping(
        EventSourceArn="arn:aws:sqs:us-east-1:952897923483:{}.fifo".format(index),
        FunctionName="Consumer-{}".format(index),
        Enabled=True,
        BatchSize=10,
    )