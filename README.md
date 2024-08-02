# Slis-Pyrra

## Intro

This is a simple CLI written in golang, where the idea is: 

- Get from Backstage the application metadata to use them as labels for pyrra
- Get from Prometheus the average time response to create Latency SLIs
- Parse all those information to create SLOs where pyrra can consume to apply in kubernetes
