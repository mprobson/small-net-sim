
BigSim
BigNetSim
TraceR
Hatchet

node(s)
- cores
  - compute speed (skip GPU for now)
- network speed GB/bps
- num. links
- topo(?)

trace
- num messages
  - size(s)
  - destination(?)
- tasks
  - num. threads
- graph/dependecy information

- formats?
  - projections? Big(Net)Sim?
  - TraceR/hatchet
  - MPI?

comm. model - A/B (future: LogP, LogGP, etc)
- latency
- bandwidth

comp. model
- overhead? for task spawn/spread

arch
- one thread? per node
- exec task, message send overhead, loop

notes/ideas
- should we use exisiting (simple?) simulator
- should we use (P)DES
