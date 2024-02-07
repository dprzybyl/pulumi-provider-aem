import * as pulumi from "@pulumi/pulumi";
import * as aem from "@pulumi/aem";

const myRandomResource = new aem.Random("myRandomResource", {length: 24});
export const output = {
    value: myRandomResource.result,
};
