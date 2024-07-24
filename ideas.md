
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

future
- load node params from json file
- load trace from file
- variable link speeds

--

flow
- spawn thread per node
- each node/thread runs - what? 'trace' 'workload'
  - simulates timestep/move forward of 'function'
  - trace: list of computes (tasks) and communicates (sends)
  - what about receieves? we would know about them as DAG inputs

- global timestep min.?
- am i just re-creating DES?


make it an array for now, worry about DAG in future

let's create a simple 2 node, 2 task execution; no messages (yet)
- how do we map tasks to nodes? manual for now...
  - but should be flexible

---

Should types be capitalized?
