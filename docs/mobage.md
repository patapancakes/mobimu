# Mobage HTTP API Reference

## Call format:

```
http://api.mbga.jp/-{endpoint}/api/{call_name}?.raw=1&.gid={game_id}&.ver={game_ver}(&.ser={phone_serial_num})(&.icc={uim_card_id})(&.uid={user_id})
```
- `endpoint`: API server endpoint; may be one of the following:
    - `global`: Global APIs
    - `c{host_id}`: Game host specific APIs; `host_id` is a number.
- `call_name`: API call's name
- `game_id`: Current game applicaion's ID:
    - **Moba Daifugo:** 1002
- `game_ver`: Current application's version.
- `phone_serial_number`: Terminal's unique serial number.
- `uim_card_id`: Terminal's UIM card's ID.
- `user_id`: Current Mobage user's ID.

Call method is POST if a request body is provided with the call, GET otherwise.

## Global API Calls:

### `/Authorize`
- **Request body:**
    ```
    {log_key}={log_value}

    ```

- **Response body:**
    - App update required:
    ```
    VERUP;

    ```
    - Authentication error:
    ```
    ERR;{msg};{url}

    ```
    - Authentication successful:
    ```
    NOTE;{msg};{url};            // Each line may appear in any order, any number of times
    AD;{ad_id};
    CONF;{key};{value};
    AUTH;{user_id};{base_url};{nickname};   // Only this line is required to appear

    ```

### `/GetAdData`
- **Response body:**
    ```
    // Insert ad data here
    ```

### `/G{game_id}_GetAvatar`

- **Request body:**
    ```
    user_id={user_id}

    ```

- **Response body:**
    ```
    {attr_key};{attr_value};                // Repeat as many times as needed

    {gif_offset};{gif_length};              // Optional if no GIF data is provided
    // GIF data goes somewhere down here
    ```

## Host API Calls:

### `/G{game_id}_Connect_Room`

- **Request body:**
    ```
    lobby_id={lobby_id}
    room_no={room_no}
    ignore_lobby_max={ignore_lobby_max}         // OPTIONAL LINE

    ```
    - `ignore_lobby_max`: Set to 1 if true.

- **Response body:**
    - Refuse connection:
    ```
    BLACK_LIST

    ```
    - Accept connection:
    ```
    {room_id}
    {last_host_msg_seq}
    {room_params}

    ```

### `/G{game_id}_Create_Room`

- **Request body:**
    ```
    lobby_id={lobby_id}
    title={title}
    max_conn={max_conn}
    params={params}

    ```
    - `lobby_id`: Lobby to create room in.
    - `title`: Room's name.
    - `max_conn`: Maximum nº of allowed connections.
    - `params`: Extra parameter information. Probably game dependent.

- **Response body:**
    ```
    {room_no}
    {room_id}
    {session_id}
    {last_host_msg_seq}
    {room_params}

    ```

### `/G{game_id}_Disconnect_Room`

### `/G{game_id}_GetLobbyInfo`
- **Request body:**
    ```
    lobby_id={lobby_id}

    ```

- **Response body:**
    - App update required:
    ```
    VERUP

    ```
    - Connection successful:
    ```
    // TODO
    ```

### `/G{game_id}_GetLobbyList`

### `/G{game_id}_Sync_Room`