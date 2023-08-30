# ingestors-file-downloader

Project description here.

## Running
 ```sh
make logs
 ```

Fake Data
```json
{
  "data": {
    "reference": {
      "year": 2023,
      "month": 8,
      "day": 23
    }
  },
  "metadata": {
    "jobId": "job_id",
    "processingDate": "2023-07-25T20:05:10Z"
  }
}
```

## Diagram

```mermaid
flowchart LR

subgraph RabbitMQQueues 
  Q1(Queue 1)
  Q2(Queue 2)
  Q3(Queue 3)
  QN(Queue N)
end

subgraph App
  subgraph Consumer
    C((Consumer))
  end

  subgraph AsyncioQueue
    AQ((Queue))
  end

  subgraph Controller
    Ctrl((Controller))
  end

  subgraph Jobs
    J((Jobs))
  end

  Q1 --> C
  Q2 --> C
  Q3 --> C
  QN --> C

  C --> MsgValid(MsgValid)
  MsgValid(MsgValid) --> AQ
  Ctrl --> AQ

  Ctrl --> Obs
  Obs --> E1
  Obs --> E2
  Obs --> E3
  Obs --> EN

  E1 --> Obs
  E2 --> Obs
  E3 --> Obs
  EN --> Obs

  Obs --> Ctrl

  Ctrl --> J
end

```
