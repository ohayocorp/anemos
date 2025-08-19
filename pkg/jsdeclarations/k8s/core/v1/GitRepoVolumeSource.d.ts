// Auto generated code; DO NOT EDIT.



/**
 * Represents a volume that is populated with the contents of a git repository. Git repo volumes do not support ownership management. Git repo volumes support SELinux relabeling.
 * 
 * DEPRECATED: GitRepo is deprecated. To provision a container with a git repo, mount an EmptyDir into an InitContainer that clones the repo using git, then mount the EmptyDir into the Pod's container.
 * 
 */
export declare class GitRepoVolumeSource {
    constructor();
    constructor(spec: GitRepoVolumeSource);

	/**
     * Directory is the target directory name. Must not contain or start with '..'.  If '.' is supplied, the volume directory will be the git repository.  Otherwise, if specified, the volume will contain the git repository in the subdirectory with the given name.
     * 
     */
    directory?: string

	/**
     * Repository is the URL
     * 
     */
    repository: string

	/**
     * Revision is the commit hash for the specified revision.
     * 
     */
    revision?: string
}