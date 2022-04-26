export class ApiResponse {
    constructor(
        public code: number,
        public statusText: string,
        public message: string,
        public body: any,
    ) {}
}