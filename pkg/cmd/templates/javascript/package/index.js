const anemos = require("@ohayocorp/anemos");

function runTest(testDirectory) {
    // Set the base output path to the test case directory so that
    // the output files are created in the correct place.
    process.env.ANEMOS_OUTPUT_PATH=`${anemos.file.currentScriptDirectory()}/${testDirectory}/output`;
    require(testDirectory);
}

runTest(`./test/basic`);