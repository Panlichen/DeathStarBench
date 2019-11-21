import pprint

direct_output = {}

key_stack = []  # stack top is the current Rightmost key
depth_stack = []  # stack top is the key's depth

for line in open("socialNetwork-topology-simple", "r"):
    if line in ['\n','\r\n']:
        break

    key = line.lstrip().replace("\n", "")
    depth = line.count(" ")
    if len(key_stack) == 0:
        key_stack.append(key)
        depth_stack.append(depth)
        continue
    
    current_depth = depth_stack[-1]
    # print(line, depth, key, current_depth)
    if depth <= current_depth:
        while depth <= depth_stack[-1]:
            key_stack.pop(); depth_stack.pop()
        current_root = key_stack[-1]
        direct_output.setdefault(current_root, []).append(key)
        key_stack.append(key)
        depth_stack.append(depth)

    elif depth > current_depth:
        direct_output.setdefault(key_stack[-1], []).append(key)
        key_stack.append(key)
        depth_stack.append(depth)

for key in direct_output:
    direct_output[key] = list(set(direct_output[key]))

print("=== out ===")
pprint.pprint(direct_output)

direct_input = {}
for key, value in direct_output.items():
    for newkey in value:
        direct_input.setdefault(newkey, []).append(key)

print("=== in ===")
pprint.pprint(direct_input)