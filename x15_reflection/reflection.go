package TwitterWalk

import "reflect"

func walk(x interface{}, fn func(input string)) {
	val := reflect.ValueOf(x)
	field := val.Field(0)
	fn(field.String())
}

// difficult to grasp at the moment
// will come back to the full implementation soon enough
