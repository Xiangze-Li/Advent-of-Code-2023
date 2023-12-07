import networkx as nx

with open("input.txt", "r") as f:
  lines = f.readlines()

G = nx.Graph()
nodes = set()

for line in lines:
  frNode, line = line.split(":")
  toNodes = line.split()

  nodes.add(frNode)
  G.add_node(frNode)

  for toNode in toNodes:
    nodes.add(toNode)
    G.add_node(toNode)
    G.add_edge(frNode, toNode, capacity=1)

nodes = list(nodes)
cut, part = nx.minimum_cut(G, nodes[0], nodes[3])

print(cut)
print(len(part[0]) * len(part[1]))
