import { View, Text, Image, TextInput, StatusBar, StyleSheet, Keyboard, TouchableWithoutFeedback } from "react-native";
import { Button } from "react-native";
import { useState } from "react";

import RegisterButton from "../components/RegisterButton";

const Login = ({navigation}) => {
    const [hidden, setHidden] = useState(true)

    return (
        <TouchableWithoutFeedback onPress={() => Keyboard.dismiss()}>
            <View style={styles.container}>
                <Image 
                    source={require("../assets/Macro.png")} 
                    style={styles.logo}
                    />
                <View style={styles.backgroundform}>
                    <Text style={styles.signupmessage}>Sign In To Start Tracking</Text>
                    <View style={styles.signup}>
                        <Text style={styles.text}>Username</Text>
                        <TextInput 
                            style={styles.input}
                            placeholder="Enter Username"
                            placeholderTextColor="#5D5D5D"
                            />
                        <Text style={styles.text}>Password</Text>
                        <TextInput 
                            style={styles.input}
                            placeholder="Enter Password" 
                            placeholderTextColor="#5D5D5D"
                            secureTextEntry={hidden}/>
                        <RegisterButton text="Log In"/>
                        <View style={styles.signedup}>
                            <Text style={styles.nosign}>Don't have an account?</Text>
                            <Button title="Sign Up" onPress={() => navigation.navigate("Signup")}/>
                        </View>
                    </View>
                </View>
            </View>
        </TouchableWithoutFeedback>
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
    signedup: {
        display: "flex",
        flexDirection: "row",
        alignItems : "center",
        justifyContent: "center"
    },
    nosign: {
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


export default Login;