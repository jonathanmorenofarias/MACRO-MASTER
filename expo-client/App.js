import {NavigationContainer } from "@react-navigation/native"
import { createStackNavigator } from "@react-navigation/stack"
import { View } from "react-native"
import Landing from "./screens/Landing"
import Login from "./screens/Login"
import Signup from "./screens/Signup"

const Stack = createStackNavigator()

export default function App() {
  return (  
    <NavigationContainer>
      <Stack.Navigator initialRouteName="Landing" screenOptions={{headerShown: false}}>
        <Stack.Screen name="Landing" component={Landing} />
        <Stack.Screen name="Login" component={Login} />
        <Stack.Screen name="Signup" component={Signup} />
      </Stack.Navigator>
    </NavigationContainer>

  )
}