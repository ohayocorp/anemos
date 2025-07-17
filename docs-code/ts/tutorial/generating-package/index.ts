import * as anemos from "@ohayocorp/anemos";

function runTest(testDirectory: string): void {
    // Change the current working directory to the test directory
    // so that the test outputs are written to the correct location.
    process.env.ANEMOS_OUTPUT_PATH=`${anemos.file.currentScriptDirectory()}/output/${testDirectory}`;
    require(testDirectory);
}

// Run basic tests.
runTest(`./test/basic`);
