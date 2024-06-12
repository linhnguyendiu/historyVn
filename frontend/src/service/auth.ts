import { message } from 'antd';
import history from '../history'
import Post from '../page/home/post'
import { error } from 'console';

export const signupRequest = async (info: any): Promise<any> => { 
    console.log('[values]', info)
    return fetch(`${process.env.REACT_APP_BASE_URL}/api/v1/users`, { 
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(info)
    })
    .then((response) => response.json())
    .then((responseData) => { 
        let {code, status, data} = responseData;
        if (code === 200) { 
            history.push('/signin')
            console.log('sign up ', responseData)
        }
    }).catch(status => { 
        throw (status)
    });
}
export const signinRequest = async (info: any): Promise<any> => { 
    console.log('[values]', info)
    return fetch(`${process.env.REACT_APP_BASE_URL}/api/v1/users/login`, { 
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(info)
    })
    .then((response) => response.json())
    .then((responseData) => { 
        let {code, status, data} = responseData;
        if (code === 200) { 
            history.push('/home')
            // console.log('sign in ', responseData)
        }
    }).catch(status => { 
        throw (status)
    });
}