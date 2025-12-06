const anemos = require("@ohayocorp/anemos");
const app = require(".");

const builder = new anemos.Builder();

app.add(builder, {
    name: "test-PACKAGE_NAME",
    namespace: "test-namespace",
    image: "test-image:v1",
});

builder.build();