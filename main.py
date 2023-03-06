import sys

import requests
import yaml


def read_Functionio() -> dict:
    """Read the FunctionIO from stdin."""
    return yaml.load(sys.stdin.read(), yaml.Loader)


def write_Functionio(Functionio: dict):
    """Write the FunctionIO to stdout and exit."""
    sys.stdout.write(yaml.dump(Functionio))
    sys.exit(0)


def result_warning(Functionio: dict, message: str):
    """Add a warning result to the supplied FunctionIO."""
    if "results" not in Functionio:
        Functionio["results"] = []
    Functionio["results"].append({"severity": "Warning", "message": message})


def main():
    """Annotate all desired composed resources with a quote from quotable.io"""
    try:
        Functionio = read_Functionio()
    except yaml.parser.ParserError as err:
        sys.stdout.write("cannot parse FunctionIO: {}\n".format(err))
        sys.exit(1)

    composite = Functionio["desired"]["composite"]["resource"]

    if "metadata" not in composite:
        print("Added metadata")
        composite["metadata"] = {}

    if "annotations" not in composite["metadata"]:
        print("Added annotations")
        composite["metadata"]["annotations"] = {}

    if not "appcat.io/effectiveVersion" in composite["metadata"]["annotations"]:
        composite["metadata"]["annotations"]["appcat.io/effectiveVersion"] = composite["spec"]["parameters"]["service"]["version"]

    write_Functionio(Functionio)


if __name__ == "__main__":
    main()
