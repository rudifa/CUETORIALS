// https://cuelang.org/docs/tutorials/tour/intro/schema/

#Conn: {
    address:  string
    port:     int
    protocol: string
}

lossy: #Conn & {
    address:  "1.2.3.4"
    port:     8888
    protocol: "udp"
}
