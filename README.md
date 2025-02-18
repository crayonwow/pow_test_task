# Proof of Work Wisdom service

### Description

A service for receiving wisdoms protected from ddos attacks using a proof of work algorithm.
To receive data, the client must receive a challenge from the server, solve it and provide it when requesting wisdom.
[hashcash](https://en.wikipedia.org/wiki/Hashcash) is used as an algorithm.

Pros:

- there is a library for go :)
- was originally created to protect against ddos attacks
  Cons:
- vulnerable to distributed work and ASIC

In real work, it is recommended to use memory-hard algorithms like [equihash](https://en.wikipedia.org/wiki/Equihash) or similar (equix).

Provided dynamic difficulty: when limits in rate limiter is hit, then difficulty increased.

used resources:
https://gitlab.torproject.org/tpo/core/torspec/-/blob/main/proposals/327-pow-over-intro.txt
https://en.wikipedia.org/wiki/Proof_of_work
https://en.wikipedia.org/wiki/Hashcash
