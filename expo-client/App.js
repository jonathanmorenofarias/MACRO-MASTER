import { NavigationContainer } from "@react-navigation/native"
import { createStackNavigator } from "@react-navigation/stack"
import { createBottomTabNavigator } from '@react-navigation/bottom-tabs';
import { getInitial } from "./components/logandsign"
import Landing from "./screens/Landing"
import Login from "./screens/Login"
import Signup from "./screens/Signup"
import Home from "./screens/Home"
import Bottom from "./components/Bottom"
import MaterialCommunityIcons from "@expo/vector-icons/MaterialCommunityIcons"
import SimpleLineIcons from "@expo/vector-icons/SimpleLineIcons"

//const Stack = createStackNavigator()
const Tab = createBottomTabNavigator()


export default function App() {
  return (  
    <NavigationContainer>
      {/* <Stack.Navigator>
        <Stack.Screen name="Landing" component={ Landing } options={{ headerShown: false }}/>
        <Stack.Screen name="Login" component={ Login } options={ () => getInitial("Login")}/>
        <Stack.Screen name="Signup" component={ Signup } options={ () => getInitial("Sign Up")}/>
      </Stack.Navigator>
       */}
      <Tab.Navigator>
        
        <Tab.Screen name="Meals" component={ Home } options={{tabBarLabel: () => <Bottom/>,tabBarIcon: () => <MaterialCommunityIcons name="food-variant" size={20} />}}/>
        <Tab.Screen name="Dashboard" component={ Home } options={{tabBarLabel: () => <Bottom/>,tabBarIcon: () => <MaterialCommunityIcons name="view-dashboard" size={20} />}}/>
        <Tab.Screen name="Plans" component={ Home } options={{tabBarLabel: () => <Bottom/>,tabBarIcon: () => <SimpleLineIcons name="notebook" size={20} />}}/>
    </Tab.Navigator>
    </NavigationContainer>

  )
}