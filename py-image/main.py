import yaml, subprocess

with open('../images.yaml', 'r') as file:
  conf = yaml.safe_load(file)

# yaml


args = ("../test/crane", "ls", "ghcr.io/open-telemetry/opentelemetry-collector-releases/opentelemetry-collector-contrib")
#Or just:
#args = "bin/bar -c somefile.xml -d text.txt -r aString -f anotherString".split()
popen = subprocess.Popen(args, stdout=subprocess.PIPE)
popen.wait()
output = popen.stdout.read().splitlines()
# print(output)

for tag in output:
  print(tag)