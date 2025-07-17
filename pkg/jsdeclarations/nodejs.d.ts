interface Console {
    log(...data: any[]): void;
}

interface Process {
    argv: string[];

    env: {
        [key: string]: string | undefined;
    };

    chdir(directory: string): void;
    cwd(): string;
}

declare var console: Console;
declare var process: Process;

declare function require(module: string): any;