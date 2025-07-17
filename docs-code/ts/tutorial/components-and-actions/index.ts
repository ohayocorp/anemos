import * as anemos from "@ohayocorp/anemos";
import { Component } from "./component";

const builder = new anemos.Builder("1.31", anemos.KubernetesDistribution.Minikube, anemos.EnvironmentType.Development);

builder.addComponent(new Component());

builder.build();