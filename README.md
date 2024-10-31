# Lightweight Kafka (with KRaft) for Railway

Single-node,persistent Kafka using Confluent's best-practice Kafka image. Uses the built-in KRaft controller for a lighter deployment that doesn't require zookeeper.

[![Deploy on Railway](https://railway.app/button.svg)](https://railway.app/template/PhV3c8?referralCode=6rOei9)

## References

- [Confluent's guide on running Kafka in KRaft mode](https://docs.confluent.io/platform/current/installation/docker/config-reference.html)
- [Configuring persistence for cp-kafka image](https://docs.confluent.io/platform/current/installation/docker/operations/external-volumes.html#data-volumes)
- [Confluent's Go Kafka SDK](https://docs.confluent.io/kafka-clients/go/current/overview.html#ak-consumer)

