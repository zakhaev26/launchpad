-module(rmq_connect).

-include_lib("amqp_client/include/amqp_client.hrl").

-compile([export_all]).

connect_to_rmq() ->
  try
    {ok, Connection} = amqp_connection:start(#amqp_params_network{
      host = "localhost",  
      port = 5672         
    }),
    io:format("Connected to RabbitMQ:~n"),

    amqp_connection:close(Connection),
    io:format("Disconnected from RabbitMQ:~n")
  catch
    error:{badarg, Reason} ->
      io:format("Connection failed: ~p~n", [Reason])
  end.
