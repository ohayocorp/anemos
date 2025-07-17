import * as anemos from "@ohayocorp/anemos";
import { Component } from "./component";
import { Options } from "./options";

export { Options };

export function add(builder: anemos.Builder, options?: Options): anemos.Component {
    const component = new Component(options);
    builder.addComponent(component);

    return component;
}