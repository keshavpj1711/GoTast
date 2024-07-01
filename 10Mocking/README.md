# Why should we know about MOCKING?

In our testing scenarios we might not always want to actually run things how we are doing in actual run,

### For example:

In our countdown example, instead of actually waiting for 3 seconds to be finished we can actually **mock** `time.Sleep` intead of **real** `time.Sleep`
So what we do is we, create a Sleeper interface with method Sleep() in it, which will have different implementations