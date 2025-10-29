/**
 * Predefined types for the environment. These are used to determine the default configurations of the components.
 * For example, in the development environment, some components may disable high availability.
 */
type EnvironmentType = string;

export const unknown: EnvironmentType;
export const development: EnvironmentType;
export const testing: EnvironmentType;
export const production: EnvironmentType;