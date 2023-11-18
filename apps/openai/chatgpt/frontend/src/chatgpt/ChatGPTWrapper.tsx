import * as React from "react";
import {useState} from "react";
import {createTheme, ThemeProvider} from "@mui/material/styles";
import Box from "@mui/material/Box";
import CssBaseline from "@mui/material/CssBaseline";
import Toolbar from "@mui/material/Toolbar";
import IconButton from "@mui/material/IconButton";
import MenuIcon from "@mui/icons-material/Menu";
import Typography from "@mui/material/Typography";
import Button from "@mui/material/Button";
import Container from "@mui/material/Container";
import ChatGPTPageText from "./ChatGPT";
import {AppBar, Card, CardContent, CircularProgress, FormControl, MenuItem, Select, Stack} from "@mui/material";
import {heraApiGateway} from "../gateway/hera";

const mdTheme = createTheme();


export function ChatGPTPage() {
    const [code, setCode] = useState('');
    const [open, setOpen] = React.useState(true);
    const toggleDrawer = () => {
        setOpen(!open);
    };
    const [language, setLanguage] = useState('plaintext');
    const handleLanguageChange = (event: any) => {
        setLanguage(event.target.value);
    };

    let buttonLabelCreate;
    let buttonDisabledCreate;
    let statusMessageCreate;
    const [requestCreateStatus, setChatRequestStatus] = useState('');
    switch (requestCreateStatus) {
        case 'pending':
            buttonLabelCreate = <CircularProgress size={20}/>;
            buttonDisabledCreate = true;
            break;
        case 'success':
            buttonLabelCreate = 'Send';
            buttonDisabledCreate = false;
            statusMessageCreate = 'Request Sent Successfully!';
            break;
        case 'insufficientTokenBalance':
            buttonLabelCreate = 'Send';
            buttonDisabledCreate = true;
            statusMessageCreate = 'Insufficient Token Balance. Email alex@zeus.fyi to request more tokens.'
            break;
        case 'error':
            buttonLabelCreate = 'Send';
            buttonDisabledCreate = false;
            statusMessageCreate = ''
            break;
        default:
            buttonLabelCreate = 'Send';
            buttonDisabledCreate = false;
            break;
    }
    const onClickSubmit = async () => {
        try {
            setChatRequestStatus('pending');
            let res: any = await heraApiGateway.sendChatGPTRequest(code)
            const statusCode = res.status;
            if (statusCode === 200 || statusCode === 204) {
                setCode(code + "\n" + "\n" + res.data)
                setChatRequestStatus('success');
            } else if (statusCode === 412) {
                setChatRequestStatus('insufficientTokenBalance');
            } else {
                setChatRequestStatus('error');
            }
        } catch (e: any) {
            if (e.response && e.response.status === 412) {
                setChatRequestStatus('insufficientTokenBalance');
            } else {
                setChatRequestStatus('error');
            }
        }
    }
    const onChange = async (textInput: string) => {
        setCode(textInput);
        // const tokenCount = await heraApiGateway.getTokenCountEstimate(textInput);
        // setTokenEstimate(tokenCount);
    };
    return (
        <ThemeProvider theme={mdTheme}>
            <Box sx={{display: 'flex'}}>
                <CssBaseline/>
                <AppBar>
                    <Toolbar
                        sx={{
                            pr: '24px', // keep right padding when drawer closed
                        }}
                    >
                        <IconButton
                            edge="start"
                            color="inherit"
                            aria-label="open drawer"
                            onClick={toggleDrawer}
                            sx={{
                                marginRight: '36px',
                                ...(open && {display: 'none'}),
                            }}
                        >
                            <MenuIcon/>
                        </IconButton>
                        <Typography
                            component="h1"
                            variant="h6"
                            color="inherit"
                            noWrap
                            sx={{flexGrow: 1}}
                        >
                            ChatGPT
                        </Typography>
                    </Toolbar>
                </AppBar>
                <Box
                    component="main"
                    sx={{
                        backgroundColor: (theme) =>
                            theme.palette.mode === 'light'
                                ? theme.palette.grey[100]
                                : theme.palette.grey[900],
                        flexGrow: 1,
                        height: '100vh',
                        overflow: 'auto',
                    }}
                >
                    <Toolbar/>
                    <Container maxWidth="xl" sx={{mt: 4, mb: 4}}>
                        <Stack direction="row" spacing={2}>
                            <Card sx={{minWidth: 250, maxWidth: 600, mt: 4, ml: 4}}>
                                <CardContent>
                                    <Typography gutterBottom variant="h5" component="div">
                                        ChatGPT Code Assistant
                                    </Typography>
                                    <Typography variant="body2" color="text.secondary">
                                        This is a code assistant using ChatGPT. It will help you prototype your code
                                        and assist with code analysis and error checking.
                                    </Typography>
                                </CardContent>
                                <Box sx={{mt: 4, ml: 4, mr: 4, display: 'flex', alignItems: 'center'}}>
                                    <Typography variant="subtitle1">Code Syntax</Typography>
                                </Box>
                                <FormControl sx={{ml: 4}}>
                                    <Select value={language} onChange={handleLanguageChange}>
                                        <MenuItem value="plaintext">PlainText</MenuItem>
                                        <MenuItem value="typescript">Typescript</MenuItem>
                                        <MenuItem value="go">Go</MenuItem>
                                        <MenuItem value="yaml">Yaml</MenuItem>
                                        <MenuItem value="pgsql">PostgreSQL</MenuItem>
                                        <MenuItem value="sql">SQL</MenuItem>
                                        <MenuItem value="json">Json</MenuItem>
                                        <MenuItem value="javascript">Javascript</MenuItem>
                                        <MenuItem value="shell">Shell</MenuItem>
                                        <MenuItem value="julia">Julia</MenuItem>
                                        <MenuItem value="python">Python</MenuItem>
                                        <MenuItem value="java">Java</MenuItem>
                                        <MenuItem value="rust">Rust</MenuItem>
                                    </Select>
                                </FormControl>
                            </Card>
                            <Box sx={{mt: 4}}>
                                <Stack direction="column" spacing={2}>
                                    <Container maxWidth="xl" sx={{mt: 4}}>
                                        {<ChatGPTPageText code={code} setCode={setCode} language={language}
                                                          onChange={onChange}/>}
                                        <Box mt={2}
                                             sx={{display: 'flex', flexDirection: 'column', alignItems: 'right'}}>
                                            <Button
                                                variant="contained"
                                                onClick={onClickSubmit}
                                                disabled={buttonDisabledCreate}
                                                sx={{backgroundColor: '#00C48C', '&:hover': {backgroundColor: '#00A678'}}}
                                            >
                                                {buttonLabelCreate}
                                            </Button>
                                            {statusMessageCreate && (
                                                <Typography variant="body2"
                                                            color={requestCreateStatus === 'error' ? 'error' : 'success'}>
                                                    {statusMessageCreate}
                                                </Typography>
                                            )}
                                        </Box>
                                    </Container>
                                </Stack>
                            </Box>
                        </Stack>
                    </Container>
                </Box>
            </Box>
        </ThemeProvider>
    );
}
