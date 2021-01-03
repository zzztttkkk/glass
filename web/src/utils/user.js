export class User {
    constructor(name, avatar, alias) {
        this.name = name || "unknown";
        this.avatar = avatar || "";
        this.alias = alias || "";
    }
}