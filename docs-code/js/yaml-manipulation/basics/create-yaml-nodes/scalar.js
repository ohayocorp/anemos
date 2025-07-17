const anemos = require("@ohayocorp/anemos");

const emptyScalar = new anemos.Scalar();
const scalarFromString = new anemos.Scalar("string");
const scalarFromNumber = new anemos.Scalar(1);
const scalarFromBoolean = new anemos.Scalar(true);
const scalarWithStyle = new anemos.Scalar("literal", anemos.YamlStyle.Literal);

const scalar = new anemos.Scalar();
// Set the value of the scalar.
scalar.setValue("string");

// Set the style of the scalar.
scalar.setStyle(anemos.YamlStyle.Literal);