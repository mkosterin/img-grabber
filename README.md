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

