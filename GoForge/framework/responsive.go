package framework

import "fmt"

func ResponsiveCSS(className, property, mobile, desktop string) string {
	return fmt.Sprintf(`
		.%s {
			%s: %s;
		}
			@media (min-width: 768px) {
				.%s {
					%s: %s;
				}
			}
	`, className, property, mobile, className, property, desktop)
}
