import { View, Text, Image, TextInput, StatusBar, StyleSheet, Linking } from "react-native";
import { Button } from "react-native";
import { useState } from "react";
import { Formik } from "formik"

import RegisterButton from "../components/RegisterButton";

const Signup = ({navigation}) => {
    const [hidden, setHidden] = useState(true)

    return (
        <View style={styles.container}>
            <Image 
                source={require("../assets/Macro.png")} 
                style={styles.logo}
                />
            <View style={styles.backgroundform}>
                <Text style={styles.signupmessage}>Sign Up To Start Tracking</Text>
                
                <Formik initialValues={{ name: '', email: '', username: '', password: '' }} 
                onSubmit={(values, actions) => {
                    fetch("/////////", {
                        method: "POST",
                        headers: {
                            "Content-Type": "application/json"  
                        },
                        body: JSON.stringify(values)
                    }).then(res => {
                        return res.json
                    }).then(data => console.log(data))
                    actions.resetForm()
                }}>
                    {(props) => (
                        <View style={styles.signup}>
                            <Text style={styles.text}>Name</Text>
                            <TextInput 
                                style={styles.input}
                                placeholder="Enter Name"
                                placeholderTextColor="#5D5D5D"
                                onChangeText={props.handleChange('name')}
                                value={props.values.name}
                                />
                            <Text style={styles.text}>Email</Text>
                            <TextInput 
                                style={styles.input}
                                placeholder="Enter Email"
                                placeholderTextColor="#5D5D5D"
                                onChangeText={props.handleChange('email')}
                                value={props.values.email}
                                />
                            <Text style={styles.text}>Username</Text>
                            <TextInput 
                                style={styles.input}
                                placeholder="Enter Username"
                                placeholderTextColor="#5D5D5D"
                                onChangeText={props.handleChange('username')}
                                value={props.values.username}
                                />
                            <Text style={styles.text}>Password</Text>
                            <TextInput 
                                style={styles.input}
                                placeholder="Enter Password" 
                                placeholderTextColor="#5D5D5D"
                                onChangeText={props.handleChange('password')}
                                value={props.values.password}
                                secureTextEntry={hidden}/>
                            <Text style={styles.text}>Re-enter Password</Text>
                            <TextInput 
                                style={styles.input}
                                placeholder="Re-enter Password"
                                placeholderTextColor="#5D5D5D" 
                                secureTextEntry={hidden}/>
                            <RegisterButton text="Sign Up" press={props.handleSubmit}/>

                            <View style={styles.loggedin}>
                                <Text style={styles.login}>Already have an account?</Text>
                                <Button title="Login" onPress={() => navigation.navigate("Login")} />
                            </View>
                        </View>
                    )}
                    

                </Formik>
                
            </View>
        </View>
    )

}

const styles = StyleSheet.create({ 
    container: {
        flex: 1,
        justifyContent: "center",
        alignItems: "center",
        backgroundColor: "black",
        gap: 30
    },
    logo: {
        height: 100,
        objectFit: "contain"
    },
    signupmessage: {
        color: "white",
        fontWeight: "bold",
        fontSize: 24
    },
    backgroundform: {
        display: "flex",
        justifyContent: "center",
        alignItems: "center",
        gap: 20,
        backgroundColor: "#353535",
        width: "90%",
        height: "70%",
        borderRadius: 10
    },
    text: {
        color: "white"
    },
    signup: {
        display: "flex",
        flexDirection: "column",
        gap: 10
    },
    loggedin: {
        display: "flex",
        flexDirection: "row",
        alignItems : "center",
        justifyContent: "center"
    },
    login: {
        textAlign: "center",
        color: "white"
    },
    input: {
        borderWidth: 1,
        borderRadius: 10,
        borderColor: "#434343",
        fontSize: 15,
        color: "white",
        backgroundColor: "#1d1d1d",
        padding: 10,
        width: 250
    }, 
    button: {
        borderWidth: .25,
        backgroundColor: "red",
        color: "red"
    }
})


export default Signup;