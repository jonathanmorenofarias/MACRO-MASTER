import {Image, View, StyleSheet} from "react-native"
import LandingButton from "../components/LandingButton"

const Landing = ({navigation}) => {
    return (
        <View style={styles.background}>
            <Image source={require("../assets/Macro.png")} style={styles.logo}/>    
            <View>
                <LandingButton 
                    text="Login" 
                    backgroundColor="#015CB8" 
                    navigate={() => navigation.navigate("Login")}/>
                <LandingButton 
                    text="Sign Up" 
                    backgroundColor="#5D5D5D"
                    navigate={() => navigation.navigate("Signup")}
                     />
            </View>
        </View>
    )
}

const styles = StyleSheet.create({
    background: {
        flex: 1,
        backgroundColor: "black",
        alignItems: "center",
        justifyContent: "space-around"
    },
    logo: {
        height: 175,
        objectFit: "contain"
    }   
})

export default Landing