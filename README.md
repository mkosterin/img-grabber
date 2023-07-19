## Tool for collecting all images from helm manifest
###
Usage:
```bash
cat ../helm_manifest.yaml | ./img-grabber
```
or
```bash
helm template <release-name> <chart-name> | ./img-grabber
```

or
```bash
img-grabber -file <manifest_file_name.yaml> -sort <asc|desc>
```

or
```bash
img-grabber -file <manifest_file_name.yaml> -mode stat
```

All possible keys:
```bash
img-grabber -h
```
