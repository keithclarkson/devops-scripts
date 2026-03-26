const fs = require('fs');
const path = require('path');
const yaml = require('js-yaml');

class Parser {
    constructor(filePath) {
        this.filePath = path.resolve(process.cwd(), filePath);
    }

    parse() {
        if (!fs.existsSync(this.filePath)) {
            throw new Error(`File not found: ${this.filePath}`);
        }

        const fileContent = fs.readFileSync(this.filePath, 'utf8');
        const extension = path.extname(this.filePath).toLowerCase();

        try {
            if (extension === '.json') {
                return JSON.parse(fileContent);
            } else if (extension === '.yaml' || extension === '.yml') {
                return yaml.load(fileContent);
            } else {
                throw new Error(`Unsupported file format: ${extension}`);
            }
        } catch (error) {
            throw new Error(`Failed to parse file: ${error.message}`);
        }
    }
}

module.exports = Parser;