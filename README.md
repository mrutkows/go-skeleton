# sbom-utility


Initially, we want to validate SPDX or CycloneDX SBOMs (JSON format only) to current standard schema.

Next, we want to parse SPDX 2.2 using a dedicated schema parser with the goal of being able to losslessly convert it to the most current CycloneDX schema.

### References

- https://github.com/spdx
- https://tools.spdx.org/app/convert/ - Used this to convert from tv format to json
    - NOTE: tool could not convert `example6-bin.spdx`; resulted in an error