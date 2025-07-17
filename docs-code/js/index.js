const anemos = require("@ohayocorp/anemos");
const thisScriptDirectory = anemos.file.currentScriptDirectory();

function runCase(path) {
    // Set the base output path to the test case directory so that
    // the output files are created in the correct place.
    process.env.ANEMOS_OUTPUT_PATH=`${thisScriptDirectory}/${path}/output`;
    require(`./${path}`);
}

runCase("yaml-manipulation/basics/create-yaml-nodes");
runCase("yaml-manipulation/basics/create-document");
runCase("yaml-manipulation/basics/modify-document");

runCase("execution-order");
runCase("helm-interop");

runCase("tutorial/generating-kubernetes-manifests");
runCase("tutorial/components-and-actions");
runCase("tutorial/components-and-actions/componentEmpty");
runCase("tutorial/component-options");
runCase("tutorial/source-control");
runCase("tutorial/generating-package");
runCase("tutorial/package-management");