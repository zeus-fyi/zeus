import axios from "axios";

export const heraApi = axios.create({
    baseURL: "http://localhost:9000"
});

class HeraApiGateway {
    async sendChatGPTRequest(prompt: string): Promise<any> {
        const url = `/v1/ui/openai/codegen`;
        try {
            // add this for setting up your own auth
            // let config = {
            //     headers: {
            //         'Authorization': `Bearer ${bearer}`
            //     },
            //     withCredentials: true,
            // }
            return heraApi.post(url, {
                prompt: prompt,
            })
        } catch (exc) {
            console.error('error sending prompt request');
            console.error(exc);
            return
        }
    }
}

export const heraApiGateway = new HeraApiGateway();

